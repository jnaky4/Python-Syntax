/*
Queue
The Queue interface defined within the Java Collections Framework defines a Collection of ordered elements that
supports element insertion at the tail and element retrieval from the head.
A LinkedList is one of several types that implements the Queue interface. A LinkedList implementation of a Queue can be
declared and created as Queue<T> queue = new LinkedList<T>(); where T represents the element's type, such as Integer or
String. Java supports automatic conversion of an object, e.g., LinkedList, to a reference variable of an interface
type, e.g., Queue, as long as the object implements the interface.
The statements import java.util.LinkedList; and import java.util.Queue; enable use of a LinkedList and Queue within a
program.
A Queue's add() method adds an element to the tail of the queue and increases the queue's size by one. A Queue's
remove() method removes and returns the element at the head of the queue. If the queue is empty, remove()
throws an exception.
 */
/*
Deque
The Deque (pronounced "deck") interface defined within the Java Collections Framework defines a Collection of ordered
elements that supports element insertion and removal at both ends (i.e., at the head and tail of the deque).
A LinkedList is one of several types that implements the Deque interface. A LinkedList implementation of a Deque can be
declared and created as Deque<T> deque = new LinkedList<T>(); where T represents the element's type, such as Integer or
String. Java supports automatic conversion of an object, e.g., LinkedList, to a reference variable of an interface
type, e.g., Deque, as long as the object implements the interface.
The statements import java.util.LinkedList; and import java.util.Deque; enable use of a LinkedList and Deque within a
program.
Deque's addFirst() and removeFirst() methods allow a Deque to be used as a stack. A stack is an ADT in which elements
are only added or removed from the top of a stack. Deque's addFirst() method adds an element at the head of the deque
and increases the deque's size by one. The addFirst() method shifts elements in the deque to make room for the new
element. The removeFirst() method removes and returns the element at the head of the deque. If the deque is empty,
removeFirst() throws an exception.
 */

import java.util.Deque;
import java.util.LinkedList;
import java.util.Queue;

public class Queues {
    public static void main( String [] args){
        Queue<String> waitList = new LinkedList<String>();

        waitList.add("Neumann party of 1");
        waitList.add("Amdahl party of 2");
        waitList.add("Flynn party of 4");

        System.out.println("Serving: " + waitList.remove());
        System.out.println("Serving: " + waitList.remove());
        System.out.println("Serving: " + waitList.remove());

        Deque<String> tripRoute = new LinkedList<String>();

        System.out.println("Depart Tokyo");
        tripRoute.addFirst("Tokyo");

        System.out.println("Transfer at Osaka");
        tripRoute.addFirst("Osaka");

        System.out.println("Arrive in Nara");
        tripRoute.addFirst("Nara");

        System.out.println("\nReturn trip: ");
        System.out.println("Depart " + tripRoute.removeFirst());
        System.out.println("Transfer at " + tripRoute.removeFirst());
        System.out.println("Arrive in " + tripRoute.removeFirst());

        //add
        //remove
        //pollFirst
        //get
        //peek



    }
}
