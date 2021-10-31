/*
A programmer may wish to lookup values or elements based on another value, such as looking up an employee's record
based on an employee ID. The Map interface within the Java Collections Framework defines a Collection that
associates (or maps) keys to values. The statement import java.util.HashMap; enables use of a HashMap.
The HashMap type is an ADT implemented as a generic class (discussed elsewhere) that supports different types of
keys and values. Generically, a HashMap can be declared and created as HashMap<K, V> hashMap = new HashMap<K, V>();
where K represents the HashMap's key type and V represents the HashMap's value type.
The put() method associates a key with the specified value. If the key does not already exist, a new entry within
the map is created. If the key already exists, the old value for the key is replaced with the new value. Thus, a
map associates at most one value for a key.
The get() method returns the value associated with a key, such as statePopulation.get("CA").
 */

    /*
    HashMap vs TreeMap
    HashMap and TreeMap are ADTs implementing the Map interface. Although both HashMap and TreeMap implement a Map, a
    programmer should select the implementation that is appropriate for the intended task. A HashMap typically provides
    faster access but does not guarantee any ordering of the keys, whereas a TreeMap maintains the ordering of keys but
    with slightly slower access. This material uses the HashMap class, but the examples above can be modified to use
    TreeMap.
     */
import java.util.Collection;
import java.util.HashMap;
import java.util.Set;

public class HashedMap {
    public static void main(String [] args){
        HashMap<String, Integer> statePopulation = new HashMap();

        // 2013 population data from census.gov
        statePopulation.put("CA", 38332521);
        statePopulation.put("AZ", 6626624);
        statePopulation.put("MA", 6692824);

        System.out.print("Population of Arizona in 2013 is ");
        System.out.print(statePopulation.get("AZ"));
        System.out.println(".");

        // 2014 estimated population
        statePopulation.put("AZ", 6871809);

        System.out.print("Population of Arizona in 2014 is ");
        System.out.print(statePopulation.get("AZ"));
        System.out.println(".");

        //putIfAbsent
        statePopulation.putIfAbsent("MO", 6137000);
        //containsKey
        boolean containsMo = statePopulation.containsKey("MO");
        //containsValue
        boolean containsVal = statePopulation.containsValue(6137000);
        //remove
        statePopulation.remove("AZ");
        
        //keySet
        Set<String> keys = statePopulation.keySet();
        //values
        Collection<Integer> values = statePopulation.values();
        //clear
        statePopulation.clear();
    }
}
