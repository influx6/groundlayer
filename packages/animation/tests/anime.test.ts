import { AnimationQueue } from '../src';

describe('Markup Tests', function () {
  it('Should be able to create a Queue', function () {
    const qu = new AnimationQueue();
    expect(qu).toBeDefined();

    const frame = qu.new();
    expect(frame).toBeDefined();

    let counter: number = 0;
    frame.add((ts: number) => {
      counter++;
    });

    qu.resume();
    expect(setTimeout).toHaveBeenCalledTimes(1);

    jest.runOnlyPendingTimers();

    expect(counter).toBe(1);
  });
});
