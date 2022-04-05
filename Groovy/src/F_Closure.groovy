class F_Closure {
    static void main(String[] args) {

        /*

            A closure definition follows this syntax:
            { [closureParameters -> ] statements }
                closureParameters are optional
                if it has a parameter, requires ->
                closures return the last line in the block
         */
        def callback = { println 'Done!' } // can use def
        assert callback instanceof Closure
        callback()


        Closure printI = { i -> println "i is $i" }
        printI(2)

        //specify return <Boolean>
        Closure<Boolean> isTextFile = { File it -> // specify Parameter Type File
            it.name.endsWith('.txt')
        }

        //variable arguments
        def concat1 = { String... it -> it.join('') }
        assert concat1('abc', 'def') == 'abcdef'



    }

}
class Person {
    public final String name
    Person(String name) { this.name = name}
    String getName() { "Name: $name" }
}

