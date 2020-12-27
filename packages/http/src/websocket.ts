/**
 * Exponent defines an interface that provides a new duration
 * for a previously provided value.
 */
export interface Exponent {
  (lastTime: number): number;
}

// SocketReader defines a interface exposes the data reader
// which is used by a Socket to read and write data to a websocket
// connection.
export interface SocketReader {
  Exhausted(socket: Socket);
  Closed(socket: Socket);
  Message(message: Event, socket: Socket);
  Errored(message: Event, socket: Socket);
  Connected(event: Event, socket: Socket);
  Disconnected(event: Event, socket: Socket);
  Reconnecting(event: Event, socket: Socket);
}

// SocketData defines a data-type for representing types of accepted data
// to write to a socket.
export type SocketData = string | ArrayBuffer | Blob | ArrayBufferView | SharedArrayBuffer;

// OneSecond represent a second in milliseconds.
export const OneSecond: number = 1000;

// OneSecond represent a 1 minute in milliseconds.
export const OneMinute: number = OneSecond * 60;

/**
 * Socket wraps a websocket connection socket, providing retries and
 * reconnection for a giving connection.
 */
export class Socket {
  public addr: string;
  public reader: SocketReader;

  private userClosed: boolean;
  private disconnected: boolean;
  private readonly maxWait: number;
  private readonly maxReconnect: number;
  private readonly exponent: Exponent | null;
  private readonly writeBuffer: Array<SocketData>;

  private lastWait: number;
  private attemptedConnects: number;
  private socket: WebSocket | null;

  /**
   *
   * @param addr Address string to connect to.
   * @param reader SocketReader to use for operations.
   * @param exponent function to use for generating exponential/linear retry times.
   * @param maxReconnects max allowed reconnection attempts.
   * @param maxWait max allowed exponential wait time in millisecond.
   */
  constructor(
    addr: string,
    reader: SocketReader,
    exponent: Exponent | null,
    maxReconnects: number | 10,
    maxWait: number | 60000,
  ) {
    this.addr = addr;
    this.socket = null;
    this.reader = reader;
    this.maxWait = maxWait;
    this.userClosed = false;
    this.exponent = exponent;
    this.disconnected = false;
    this.attemptedConnects = 0;
    this.lastWait = OneSecond;
    this.maxReconnect = maxReconnects;
    // this.readBuffer = new Array<Event>();
    this.writeBuffer = new Array<SocketData>();
  }

  connect() {
    if (this.socket) {
      return;
    }

    if (this.attemptedConnects >= this.maxReconnect) {
      this.reader.Exhausted(this);
      return;
    }

    const socket = new WebSocket(this.addr);
    socket.addEventListener('open', this._opened.bind(this));
    socket.addEventListener('error', this._errored.bind(this));
    socket.addEventListener('message', this._messaged.bind(this));
    socket.addEventListener('close', this._disconnected.bind(this));
    this.socket = socket;
    this.disconnected = false;
  }

  send(message: SocketData) {
    if (this.disconnected) {
      this.writeBuffer.push(message);
      return;
    }
    this.socket!.send(message);
  }

  // reset the socket to allow reconnection to address.
  reset(): void {
    this.attemptedConnects = 0;
    this.lastWait = OneSecond;
  }

  end(): void {
    this.userClosed = true;
    this.reader.Closed(this);
    this.socket!.close();
    this.socket = null;
  }

  _disconnected(event: Event): void {
    this.reader.Disconnected(event, this);

    this.disconnected = true;
    this.socket = null;

    // if it was a user-induced closure, then return;
    if (this.userClosed) {
      return;
    }

    let nextWait = this.lastWait;
    if (this.exponent) {
      nextWait = this.exponent(nextWait);
      if (nextWait > this.maxWait) {
        nextWait = this.maxWait;
      }
    }

    setTimeout(this.connect.bind(this), nextWait);
    this.attemptedConnects++;
  }

  _opened(event: Event) {
    this.reader.Connected(event, this);
    // if we've buffered up messages, then write them out immediately
    // to the socket.
    while (this.writeBuffer.length > 0) {
      const message = this.writeBuffer.shift();
      this.socket!.send(message!);
    }
  }

  _errored(event: Event) {
    this.reader.Errored(event, this);
  }

  _messaged(event: Event) {
    this.reader.Message(event, this);
  }
}
