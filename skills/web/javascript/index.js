

import pkg from './profile.js';

console.log(pkg.firstname);
console.log(pkg.lastname);
console.log(pkg.age);
console.log(pkg.city);
console.log(pkg.country);
console.log(pkg.email);
console.log(pkg.phone);
console.log(pkg.address);
console.log("================================================");

var o = {
    name: 'John',
    age: 20,
    city: 'New York',
    country: 'USA',
    email: 'john@example.com',
    phone: '1234567890',
    address: '123 Main St, Anytown, USA',
}

for (var key in o) {
    console.log(key,o[key]);
}
console.log("================================================");

var checkNumber = new Promise(function(success, failed) {
    setTimeout(function() {
      var number = Math.random() * 100; // 随机生成 0-100 的数字
      if (number > 50) {
        success("成功：数字 " + number + " 大于 50");
      } else {
        failed("失败：数字 " + number + " 太小");
      }
    }, 1000); // 1 秒后执行
  });


checkNumber
    .then(function(data) {
        console.log(data);
    })
    .catch(function(error) {
        console.log(error);
    });
// console.log("===============hahah=================================");