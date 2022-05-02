

{
    //block scope, 
    //if and for are block scope
    var result = 4;
} 

//use let or const instead
for(var i = 0; i < 10; i++){
    
}
i = 10; // bad var scoping

{{{ 
    var v = 42
    let u = 10
}}} //nested block scopes
v 
//u wont work

function sum(a,b){ 
    var result = a+b //function scope, not reachable
}