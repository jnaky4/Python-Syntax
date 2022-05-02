
const X = function (){
    //this keyword is the caller of X function
}
const y = () => {
    //this keyword is function scope
    
}
//regular func give access to their calling environment
//arrow func give access to their defining environment

const testerObj = {
    func1: function() { //regular function
        console.log('func1', this);
    },
    func2: () => { // arrow function
        console.log('func2', this);
    }
}
console.log(this) //references env
testerObj.func1(); //references testerObj
testerObj.func2(); //references env

//shorthand for single line function
const square = (a) => a * a;
//same as 
const square1 = (a) => {
    return a * a;
};
//useful for array functions
[1,2,3,4].Map(a => a * a)





