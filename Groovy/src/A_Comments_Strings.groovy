#!/usr/bin/env groovy
/*
    Beside the single-line comment, there is a special line comment, often called the shebang line understood
    by UNIX systems which allows scripts to be run directly from the command-line, provided you have installed
    the Groovy distribution and the groovy command is available on the PATH.
 */
import java.time.LocalDateTime


class A_Comments_Strings {
    static void main(String[] args) {
        // Single line comment
        /*
            Multi Line
            Comment
         */

        /**@
         * A Groovy Doc String
         * Text literals are represented in the form of a chain of characters called strings.
         * Groovy lets you instantiate java.lang.String objects, as well as GStrings (groovy.lang.GString)
         * which are also called interpolated strings in other programming languages.
         */


        /*
        Double-quoted strings: plain java.lang.String if there’s no interpolated expression,
        but are groovy.lang.GString instances if interpolation is present.
         */
        println "hello"

        /*
            string interpolation: groovy.lang.GString
            Interpolation is the act of replacing a placeholder in the string with its value upon evaluation of the string.
         */
        def example = "print me"
        println "hello ${example}"
        println "hello $example"

        //Triple-single-quoted strings are plain java.lang.String and don’t support interpolation.
        println '''triple - quote string ${example} - doesn\'t work'''

        def first = "john"
        def last = "doe"
        def list = [1,2,3,4]
        /*
            Triple-quoted strings may span multiple lines. The content of the string can cross line boundaries
            without the need to split the string in several pieces and without
            concatenation or newline escape characters:
            \ skips the new line created
         */
        def quotes = """\
    first name: $first
    last name:  $last

    this will print as is
    $list
    ${LocalDateTime.now()}
    ${2 + 4}
        """
//        println quotes

        def unicode_example = 'The Euro currency symbol: \u20AC' //unicode escape sequence
//        println unicode_example

        def person = [name: 'Guillaume', age: 36]
//        println "$person.name is $person.age years old" //dotted expressions

        /*
            Warning!!
            GString and Strings having different hashCode values, using GString as Map keys should be avoided,
            especially if we try to retrieve an associated value with a String instead of a GString.
         */
        assert "one: ${1}".hashCode() != "one: 1".hashCode()

        //can cause errors
        try {
            def key = "a"
            //def m = ["$key": "a"] bad way
            def m = [(key): "a"]
            println m["a"]
        }
        catch(Exception e){
            println e
        }

        /*
            Slashy String:
            Groovy offers slashy strings, which use / as the opening and closing delimiter. Slashy strings are
            particularly useful for defining regular expressions and patterns, as there is no need to escape backslashes.
         */
        def fooPattern = /.*foo.*/
        def escapeSlash = /The character \/ is a forward slash/
        def color = "red"
        def interpolatedSlashy = /a ${color} car/

        /*
            Dollar Slashy String:
            Dollar slashy strings are multiline GStrings delimited with an opening $/ and a closing /$.
            The escaping character is the dollar sign, and it can escape another dollar, or a forward slash.
         */
        def name = "Guillaume"
        def date = "April, 1st"

        def dollarSlashy = $/
            Hello $name,
            today we're ${date}.
        
            $ dollar sign
            $$ escaped dollar sign
            \ backslash
            / forward slash
            $/ escaped forward slash
            $$$/ escaped opening dollar slashy
            $/$$ escaped closing dollar slashy
        /$

        assert [
                'Guillaume',
                'April, 1st',
                '$ dollar sign',
                '$ escaped dollar sign',
                '\\ backslash',
                '/ forward slash',
                '/ escaped forward slash',
                '$/ escaped opening dollar slashy',
                '/$ escaped closing dollar slashy'
        ].every { dollarSlashy.contains(it) }

        def string_table =  """\
String name             | String syntax | Interpolated | Multiline | Escape character
Single-quoted           | '…​'                                       \\
Triple-single-quoted    | '''…​'''                                   \\
Double-quoted           | "…​"                                       \\
Triple-double-quoted    | ${""" ...  """}                            \\
Slashy                  | /…​/                                       \\
Dollar slashy           | \$/…​/\$                                   \$
        """
//        println string_table

    }

}
