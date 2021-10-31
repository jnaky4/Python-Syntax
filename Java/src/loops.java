import java.util.ArrayList;

public class loops {
    public static void main(String [] args) {


        ArrayList<String> teamRoster = new ArrayList<String>();
        String playerName;
        int i;

// Adding player names
        teamRoster.add("Mike");
        teamRoster.add("Scottie");
        teamRoster.add("Toni");

        System.out.println("Current roster: ");

//       basic loop
        for (i = 0; i < teamRoster.size(); ++i) {
            playerName = teamRoster.get(i);
            System.out.println(playerName);
        }


//      enhanced for loop
        for (String Name : teamRoster) {
            System.out.println(Name);
        }
    }
}
