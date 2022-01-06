/****** 买卖股票的最佳时机 ******/ 
const maxProfit = function(prices) {
  if(prices.length<2) return 0
  let m = 0;
  for (let i=1;i<prices.length;i++) {
    m += Math.max(0,prices[i] - prices[i-1])
  }
  return m
};
console.log(maxProfit([7,2,3,1,5,6])); // ==> 6