class D_Lists {
    static void main(String[] args) {
//        syntax()
//        listIteration()
//        listFiltering()
//        sets()

    }
    static void syntax(){
        //        Groovy doesn’t define its own collection classes. The list is an instance of Java’s java.util.List interface
        def heterogeneous = [1, "a", true]
        /*
            possible to use a different backing type for our lists,
            using type coercion with the as operator, or with explicit type declaration for your variables:
         */
        def arrayList = [1, 2, 3] // default type ArrayList
        assert arrayList instanceof ArrayList // default

        int[] explicit_arrayList = [1, 2, 3]  // explicit type declaration
        def linkedList = [2, 3, 4] as LinkedList // using as
        LinkedList otherLinked = [3, 4, 5] // explicit type declaration

        def three = arrayList[-1] // negative indexing, last element of list
        linkedList << three  // use the << left shift operator to append an element at the end of the list

        def primes = new int[] {2, 3, 5, 7, 11} //java style array initialization
    }

    static void listIteration(){
        def list = [2, 3, 12, 5, 12, 5345, 1, 354, 21, 3245, 6, 34, 2, 345, 234, 12, 24323, 5, 2, 23, 6, 3]
        list = list.each { it * 2 } // does nothing, returns the original collection
        println list
        list = list.collect{ it * 2 } // does work
        println list

        list.eachWithIndex{ int entry, int i -> print "index: $i is $entry, "}
    }

    static void listFiltering() {
        def list = [2, 3, 12, 5, 12, 5345, 1, 354, 21, 3245, 6, 34, 2, 345, 234, 12, 24323, 5, 2, 23, 6, 3]
        def alphabeta = ['a', 'b', 'c', 'd']
        println list.min()
        println alphabeta.max()
        println list.find { it > 1 && it < 200 } // only finds first
        println list.findAll { it > 1 && it < 200 }
        println list.findIndexOf { it == 2} // first instance only
        println list.findLastIndexOf { it == 2 }
        println list.sum{ it < 10 ? it : 0} // sum with condition < 10
        //custom sum
        println alphabeta.sum{ it == 'a' ? 1 : it == 'b' ? 2 : it == 'c' ? 3 : it == 'd' ? 4 : it == 'e' ? 5 : 0 }
        println list.join("->")
        println list.inject("concatenate: "){str, item -> str + item }
        println list.collect{ it * 2 }
        println list.sort()
        println list.count{ it % 2 == 0}

        def nested = [1, 2, 3, [4, 5, [6, 7, 8]]]
        println nested.flatten()
        println nested.collectNested { it % 2 == 0 ? it : 0 }

        //todo examples
//        println list.eachParallel {}
//        println list.parallelStream()
//        println list.spliterator()
//        println list.asConcurrent {}
//        println list.asCollection()
//        println list.asSynchronized()
//        println list.collectParallel {}
//        list.collate()
//        list.stream()
//        list.grep()
//        list.grepParallel()
//        list.groupBy()
//        list.groupByParallel {}
//        list.eachCombination {}
//        list.eachPermutation {}
//        list.makeConcurrent()
//        list.transpose()
//        list.dump()
//        list.metaClass {}
//        list.macro {}
//        list.finalize()
//        list.notify()

        //todo closure example of timeout
        //todo example of object map

    }
    

    static void sets(){
        def list = [2, 3, 12, 5, 12, 5345, 1, 354, 21, 3245, 6, 34, 2, 345, 234, 12, 24323, 5, 2, 23, 6, 3]
        def list2 = [2, 5, 21, 6, 0]
        println list.contains(2)
        println 2 in list
        println list.containsAll(2,4)
        println list.intersect(list2) // joins the list of only matching values
        println list.disjoint(list2) // returns boolean if disjointed
    }
}
