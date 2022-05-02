//destructuring
//works for arrays and objects

//const PI = Math.PI;
//const E = Math.E;
//const SQRT2 =  Math.SQRT2;

//same as
const {PI, E, SQRT2} = Math;

// typical imports
// const {Component, Fragment, useState} = require('react');


//object parameter destructuring
const circle = {
    label: 'circleX',
    radius: 2,
};

//default parameters with destructuring
const circleArea = ({radius}, {precision = 2} = {}) => //if you pass object to function, destructuring can reference class members, ie radius
    (PI * radius * radius).toFixed(precision); 

console.log(circleArea(circle)); //here we passe the circle object, so destructuring works

//destructuring array values
const [fName, lName,,age] = ["Jacob", "Alongi", "Male", 30];
console.log(fName);
const [first, ...restOfItems] = [10, 20, 30, 40]; //destructure first, restOperator ... : restOfItems is new Array
console.log(restOfItems); 

const obj = {
    temp1: '001',
    temp2: '002',
    firstName: "John",
    lastName: "Doe",
};

const {temp1, temp2, ...person} = data; //can create new object from parts of old

const copyOfRest = [...restOfItems]; //copying arrays



const newObject = { //shallow copies
    ...person
};
