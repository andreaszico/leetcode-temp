/**
 * @param {number} n
 * @return {number}
 */

function fibo(n) {
  if (n <= 1) {
    return 1;
  }

  return fibo(n - 1) + fibo(n - 2);
}

var climbStairs = function (n) {
  return fibo(n);
};

console.log(climbStairs());
