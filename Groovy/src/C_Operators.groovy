class C_Operators {
    static void main(String[] args) {
        //elvis operator
        def user = ['name': 'jake']
        def displayName = user.name ?: 'Anonymous' //short hand ternary
        def displayNickname = user.name ?= 'Anonymous' //same as above

        //safe navigation operator
        def person = user.find { it.value == "Jack" }
        def name = person?.value //avoids null pointer
        println name == null

        //ternary operator
        def string = "hello"
        def result = ""
        //if else same as ternary below
        if (string!=null && string.length()>0) {
            result = 'Found'
        } else {
            result = 'Not found'
        }
        // assignment = condition ? if True do : else do
        result = (string != null && string.length() > 0 ) ? 'Found' : 'Not found'
        result = string ? 'Found' : 'Not found' // Using groovy truth

        //method pointer operator


        //direct field access operator
        def useless = new Employee("Bob")
        assert useless.name == 'Name: Bob'
        assert useless.@name == 'Bob' // use of .@ forces usage of the field instead of the getter
    }
}
class Employee {
    public final String name
    Employee(String name) { this.name = name}
    String getName() { "Name: $name" }
}