/***** 存在重复的元素 *****/ 
const containsDuplicate = function(nums) {
  if(nums.length<2)return false
  nums.sort();
  let i = 1;
  let result = false;
  while (i <= nums.length) {
    if(nums[i] == nums[i-1]){
      result = true;
      break
    }else{
      i++
    }
  }
  return result
};
console.log(containsDuplicate([2,14,18,22,22]));