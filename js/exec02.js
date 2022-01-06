/****** 旋转数组 ******/ 
// const rotateArr = function(nums,k) {  // 超时
//   let len = nums.length;
//   k %= len;
//   if(len < 2 || k == 0) return nums
//   for (let i = 0; i < k; i++) {
//     nums.unshift(nums[len - 1])
//   }
//   nums.splice(len,k);
//   return nums
// }
// console.log(rotateArr([1,2,3,4,5,6,7],3)); // ==> [6,7,1,2,3,4,5]

const rotateArr = function(nums,k) {
  const n = nums.length
  let res = new Array(n)
  for(let i = 0; i < n; i++){
    res[(i+k)%n] = nums[i]
  }
  for(let i = 0;i < n; i++){
    nums[i] = res[i]
  }
  return nums
}
console.log(rotateArr([1,2,3,4,5,6,7],2)); // ==> [6,7,1,2,3,4,5]