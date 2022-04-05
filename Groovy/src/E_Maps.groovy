class E_Maps {
    static void main(String[] args) {
        def colors = [red: '#FF0000', green: '#00FF00', blue: '#0000FF'] // map/dictionary/associative array
        println colors['red'] // subscript notation
        println colors.red // property notation

        assert colors instanceof LinkedHashMap
        println colors.containsKey('red')

        /*
            When you need to pass variable values as keys in your map definitions,
            you must surround the variable or expression with parentheses:
         */
        def key = 'jake'
        def person = [(key): '30']

    }
}
