export function LinearSum(p0, p1, t) {
  return (p1 - p0) * t + p0;
}

export function BernsteinDivision(n, i) {
  const fc = FactorialGenerator();
  return fc(n) / fc(i) / fc(n - i);
}

export function FactorialGenerator() {
  const a = [1];
  return function (n) {
    let s = 1;
    if (a[n]) {
      return a[n];
    }

    for (let i = n; i > 1; i--) {
      s *= i;
    }

    a[n] = s;
    return s;
  };
}

export function CatmullRomSum(p0, p1, p2, p3, t) {
  var v0 = (p2 - p0) * 0.5;
  var v1 = (p3 - p1) * 0.5;
  var t2 = t * t;
  var t3 = t * t2;
  return (2 * p1 - 2 * p2 + v0 + v1) * t3 + (-3 * p1 + 3 * p2 - 2 * v0 - v1) * t2 + v0 * t + p1;
}

export function Linear(v, k) {
  const m = v.length - 1;
  const f = m * k;
  const i = Math.floor(f);

  if (k < 0) {
    return LinearSum(v[0], v[1], f);
  }
  if (k > 1) {
    return LinearSum(v[m], v[m - 1], m - f);
  }
  return LinearSum(v[i], v[i + 1 > m ? m : i + 1], f - i);
}

export function Bezier(v, k) {
  const n = v.length - 1;
  const pw = Math.pow;

  let b = 0;
  for (let i = 0; i <= n; i++) {
    b += pw(1 - k, n - i) * pw(k, i) * v[i] * BernsteinDivision(n, i);
  }
  return b;
}

export function CatmullRom(v, k) {
  const m = v.length - 1;
  let f = m * k;

  let i = Math.floor(f);
  if (v[0] === v[m]) {
    if (k < 0) {
      f = m * (1 + k);
      i = Math.floor(f);
    }
    return CatmullRomSum(v[(i - 1 + m) % m], v[i], v[(i + 1) % m], v[(i + 2) % m], f - i);
  }

  if (k < 0) {
    return v[0] - (CatmullRomSum(v[0], v[0], v[1], v[1], -f) - v[0]);
  }

  if (k > 1) {
    return v[m] - (CatmullRomSum(v[m], v[m], v[m - 1], v[m - 1], f - m) - v[m]);
  }
  return CatmullRomSum(v[i ? i - 1 : 0], v[i], v[m < i + 1 ? m : i + 1], v[m < i + 2 ? m : i + 2], f - i);
}
