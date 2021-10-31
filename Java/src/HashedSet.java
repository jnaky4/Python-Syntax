/*
The Set interface defined within the Java Collections Framework defines a Collection of unique elements. The Set
interface supports methods for adding and removing elements, as well as querying if a set contains an element. For
example, a programmer may use a set to store employee names and use that set to determine which customers are eligible
for employee discounts.
The HashSet type is an ADT implemented as a generic class that supports different types of elements. A HashSet can be
declared and created as HashSet<T> hashSet = new HashSet<T>(); where T represents the HashSet's type, such as Integer
or String. The statement import java.util.HashSet; enables use of a HashSet within a program.
 */
import java.util.HashSet;

public class HashedSet {
    public static void main( String [] args){
        HashSet<String> ownedBooks = new HashSet<String>();

        ownedBooks.add("A Tale of Two Cities");
        ownedBooks.add("The Lord of the Rings");

        System.out.println("Contains \"A Tale of Two Cities\": " +
                ownedBooks.contains("A Tale of Two Cities"));

        ownedBooks.remove("The Lord of the Rings");

        System.out.println("Contains \"The Lord of the Rings\": " +
                ownedBooks.contains("The Lord of the Rings"));
    }
}
