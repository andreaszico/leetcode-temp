/**
 * @param {number[]} digits
 * @return {number[]}
 */
var plusOne = function (digits) {
  let last = digits.length - 1;

  while (last != -2) {
    var sum = digits[last] + 1;

    if (sum > 9) {
      digits[last] = 0;
      last--;
      continue;
    }

    if (last < 0) {
      digits.unshift(1);
      return digits;
    }

    digits[last] = sum;
    return digits;
  }
};

console.log(plusOne([1, 9, 9, 9, 9, 9, 9, 0, 9]));
console.log(plusOne([3, 8, 9]));
console.log(plusOne([9, 9]));
console.log(plusOne([9]));
