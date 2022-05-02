//object literals
const mysteryVariable = 'total';
const sum = 4 + 2;
const obj = {
    [mysteryVariable]: 10, //dynamic property syntax
    sum, //shorthand
}
console.log(obj.total);// variable is now total, not mystery.
console.log(obj.mysteryVariable);// mysteryVariable is undefined