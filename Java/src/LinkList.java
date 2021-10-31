/*
The List interface defined within the Java Collections Framework defines a Collection of ordered elements, i.e., a
sequence. The List interface supports methods for adding, modifying, and removing elements.
A LinkedList is one of several types that implement the List interface. The LinkedList type is an ADT implemented as a
generic class that supports different types of elements. A LinkedList can be declared and created as
LinkedList<T> linkedList = new LinkedList<T>(); where T represents the LinkedList's type, such as Integer or String.
The statement import java.util.LinkedList; enables use of a LinkedList within a program.
A LinkedList supports insertion of elements either at the end of the list or at a specified index. If an index is not
provided, as in authorList.add("Martin");, the add() method adds the element at the end of the list. If an index is
specified, as in authorList.add(0, "Rowling");, the element is inserted at the specified index, with the list element
previously located at the specified index appearing after the inserted element.
 */

import java.util.LinkedList;
import java.util.ListIterator;

public class LinkList {
    public static void main(String [] args) {

        LinkedList<String> authorsList = new LinkedList<>();
        String authorName;
        String upperCaseName;
        ListIterator<String> listIterator;

        //LinkedList add()
        authorsList.add("Gamow");
        authorsList.add("Greene");
        authorsList.add("Penrose");


        listIterator = authorsList.listIterator();

        while (listIterator.hasNext()) {
            authorName = listIterator.next();
            upperCaseName = authorName.toUpperCase();
            listIterator.set(upperCaseName);
        }

        listIterator = authorsList.listIterator();

        while (listIterator.hasNext()) {
            authorName = listIterator.next();
            System.out.println(authorName);
        }

        //get
        String temp = authorsList.get(2);
        String first = authorsList.getFirst();
        String last = authorsList.getLast();
        //set
        authorsList.set(2, "Taco");
        //size
        int size = authorsList.size();
        //remove
//        authorsList.removeFirst();
//        authorsList.removeLast();
        authorsList.removeFirstOccurrence("Taco");

    }

}
