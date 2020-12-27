import * as rafPolyfill from './raf-polyfill';

// FrameRequestCallback defines a callback interface.
export interface FrameRequestCallback {
  (time: number): void;
}

// AnimationQueue provides a queuing mechanism to create sectioned animations where
// one queue represent a set of an animation sequence or frames. This frames have no
// correlation with a Frame in animation but rather the idea of a frame of callbacks
// that may be connected together, basically a callback set. The idea is the queue
// sits on top of RequestAnimationFrame sequencing a set of Frames who control a set
// of callbacks.
//
// Start -> AnimationQueue.bind
// Stop -> AnimationQueue.unbind.
//
export class AnimationQueue {
  public readonly frames: Array<AFrame>;
  private skip: boolean;
  private binded: boolean;
  private requestAnimationID: number;
  private rafProvider: rafPolyfill.RAF;
  private readonly bindCycle: FrameRequestCallback;

  constructor() {
    this.skip = false;
    this.binded = false;
    this.requestAnimationID = -1;
    this.frames = new Array<AFrame>();
    this.bindCycle = this.cycle.bind(this);
    this.rafProvider = rafPolyfill.GetRAF();
  }

  // new returns a new Frame from queue.
  new(): AFrame {
    const newFrame = new AFrame(this.frames.length, this);
    this.frames.push(newFrame);
    this.bind();
    return newFrame;
  }

  // Add adds provided frame into queue.
  add(f: AFrame) {
    f.queueIndex = this.frames.length;
    f.queue = this;
    this.frames.push(f);
    this.bind();
  }

  resume() {
    this.skip = false;
    this.bind();
  }

  pause() {
    this.skip = true;
  }

  // unbind unbinds the giving queue from request animation frame.
  unbind() {
    if (!this.binded) {
      return null;
    }

    this.rafProvider.cancelAnimationFrame(this.requestAnimationID);
  }

  bind() {
    if (this.binded) return null;
    this.requestAnimationID = this.rafProvider.requestAnimationFrame(this.bindCycle, null);
    this.binded = true;
  }

  cycle(ms: number) {
    if (this.frames.length === 0) {
      this.binded = false;
      return;
    }

    // animate all frames connected with this queue.
    this.frames.forEach(function (f: AFrame) {
      if (!f.paused()) {
        f.animate(ms);
      }
    });

    this.bind();
  }
}

// AFrame provides sets of callbacks sets which will be executed at the same time
// by the underline queue.
export class AFrame {
  public queueIndex: number;
  public queue?: AnimationQueue;

  private skip: boolean;
  private readonly callbacks: Array<FrameRequestCallback>;

  constructor(index: number, queue: AnimationQueue) {
    this.skip = false;
    this.queue = queue;
    this.queueIndex = index;
    this.callbacks = new Array<FrameRequestCallback>();
  }

  // Add adds new callback into giving frame callback group.
  add(callback: FrameRequestCallback) {
    this.callbacks.push(callback);
  }

  // Clear clears all items from underline callback.
  clear() {
    this.callbacks.length = 0;
  }

  // Paused returns true/false if giving frame is paused.
  paused(): boolean {
    return this.skip;
  }

  // Pause pauses giving frame.
  pause() {
    this.skip = true;
  }

  // stop removes giving frame from it's underline queue.
  stop() {
    this.pause();

    if (this.queueIndex === -1) {
      return null;
    }

    if (this.queue!.frames.length == 0) {
      this.queue = undefined;
      this.queueIndex = -1;
      return null;
    }

    const total = this.queue!.frames.length;
    if (total == 1) {
      this.queue!.frames.pop();
      this.queue = undefined;
      this.queueIndex = -1;
      return null;
    }

    this.queue!.frames[this.queueIndex] = this.queue!.frames[total - 1];
    this.queue!.frames.length = total - 1;

    this.queue = undefined;
    this.queueIndex = -1;
  }

  animate(ts: number) {
    for (let index in this.callbacks) {
      const callback = this.callbacks[index];
      callback(ts);
    }
  }
}

// Function defines an interface used to represent a function.
export interface Function {
  (): void;
}

interface functionWrapper {
  (Function): void;
}

/**
 * ChangeManager implements a read-mutation manager which queues up calls that
 * only do reads together and executes them before executing calls that do
 * mutations together. It is meant to help provide as much effect free functions
 * and thinking in how we read and apply changes from elements as reading can
 * greatly affect UI layout thrashing when not done right.
 *
 * ChangeManager implements the base mechanism to manage and coordinate read
 * calls ensuring to execute as much read calls and write calls together
 * ensuring to minimize the effects that can come when done at the same time
 * or together in the context of UI or ever change values.
 *
 * ChangeManager uses a frame within an animation queue which relies on RAF
 * to coordinate it's execution of read and write calls.
 */
class ChangeManager {
  private readonly reads: Array<Function>;
  private readonly writes: Array<Function>;
  private readonly frame: AFrame;

  private readState: boolean;
  private inReadCall: boolean;
  private inWriteCall: boolean;
  private scheduled: boolean;

  static drainTasks(q: Array<Function>, wrapper: functionWrapper): void {
    let task: Function = q.shift()!;
    while (task) {
      if (wrapper !== null) {
        wrapper(task);
        task = q.shift()!;
        continue;
      }

      task();
      task = q.shift()!;
    }
  }

  constructor(queue: AnimationQueue) {
    this.reads = new Array<Function>();
    this.writes = new Array<Function>();
    this.readState = false;
    this.inReadCall = false;
    this.inWriteCall = false;
    this.scheduled = false;
    this.frame = queue.new();
  }

  mutate(fn: Function): void {
    this.writes.push(fn);
    this._schedule();
  }

  read(fn: Function): void {
    this.reads.push(fn);
    this._schedule();
  }

  _schedule() {
    if (this.scheduled) {
      return;
    }

    this.scheduled = true;
    this.frame.add(this._runTasks.bind(this));
  }

  /**
   * runTask contains needed logic to coordinate the
   * call to read and write.
   * @private
   */
  _runTasks() {
    // execute reads first.
    const readError = this._runReads();
    if (readError !== null && readError !== undefined) {
      // schedule remaining tasks to execute.
      this.scheduled = false;
      this._schedule();

      // throw received error.
      throw readError;
    }

    // execute writes second.
    const writeError = this._runWrites();
    if (writeError !== null && writeError !== undefined) {
      // schedule remaining tasks to execute.
      this.scheduled = false;
      this._schedule();

      // throw received error.
      throw writeError;
    }

    if (this.reads.length > 0 || this.writes.length > 0) {
      // schedule remaining tasks to execute.
      this.scheduled = false;
      this._schedule();
      return;
    }

    this.scheduled = false;
  }

  _runReads(): Error | null {
    try {
      ChangeManager.drainTasks(this.reads, this._execReads.bind(this));
    } catch (e) {
      return e;
    }
    return null;
  }

  _execReads(task: Function): void {
    this.inReadCall = true;
    task();
    this.inReadCall = false;
  }

  _runWrites(): Error | null {
    try {
      ChangeManager.drainTasks(this.writes, this._execWrite.bind(this));
    } catch (e) {
      return e;
    }
    return null;
  }

  _execWrite(task: Function): void {
    this.inWriteCall = true;
    task();
    this.inWriteCall = false;
  }
}
