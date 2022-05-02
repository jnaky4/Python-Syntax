class Person {
    constructor(name){
        this.name = name;
    }
    greet(){
        console.log(`Hello ${this.name}`);
    }
}

class Student extends Person{
    constructor(name, level){
        super(name)
        this.level = level;
    }
    greet(){
        console.log(`Hello ${this.name} from ${this.level}`);
    }
}


const o1 = new Person("Max");
const o2 = new Student("Lisa", "1st Grade");
o2.greet = () => console.log('Special!');

o1.greet();
o2.greet();
