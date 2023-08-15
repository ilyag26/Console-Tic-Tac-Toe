    import java.util.Arrays;
    import java.util.Scanner;

    public class Main {
        public static int pos1_user; //first coordinate for user position
        public static int pos2_user; //second coordinate for user position
        public static int pos1_computer; //first coordinate for computer position
        public static int pos2_computer; //second coordinate for computer position
        public static boolean game = false;
        public static void main(String[] args) {
            int[] arrBusy = new int[9];

            int step_game = 0; //counter for knowing step of game
            int step = 0; //step counter
            boolean busy = false;
            int min = 1; // Minimum value of range
            int max = 9; // Max value of range
            // field for game
            String[][] a = {
                    {"-", "-", "-"},
                    {"-", "-", "-"},
                    {"-", "-", "-"}
            };
            // print rules
            String[][] rules = {
                    {"1", "2", "3"},
                    {"4", "5", "6"},
                    {"7", "8", "9"}
            };

            Scanner scan = new Scanner(System.in);

            System.out.println("------JAVA TIC-TAC-TOE-------");
            //first print field of game
            Logic.PrintField(a);
            System.out.println("-------------");
            System.out.println("For make step just choose any number:");
            System.out.println("-------------");
            //print rules of game
            Logic.PrintRules(rules);
            System.out.println("-------------");
            while (!game) {
                System.out.println("Now make step");
                System.out.print("Enter any number: ");
                int number = scan.nextInt();
                arrBusy[step] = number;
                step += 1;
                // for making step
                Logic.Step(number);
                step_game++;
                //checking win
                Logic.CheckWin(a);
                if(game){
                    break;
                }
                //print field of game after making step
                Logic.PrintFieldUser(a,pos1_user,pos2_user);
                System.out.println("-------------");
                while (!busy) {
                    int random_int = (int) Math.floor(Math.random() * (max - min + 1) + min);
                    boolean checks = Arrays.stream(arrBusy).anyMatch(i -> i == random_int);
                    if (!checks) {
                        busy = true;
                        Logic.StepComputer(random_int);
                    }
                }
                busy = false;
                step_game++;
                //checking win
                Logic.CheckWin(a);
                if(game){
                    break;
                }
                // print field after computer's step
                Logic.PrintFieldComputer(a,pos1_computer,pos2_computer);
                if (step_game == 12) {
                    game = true;
                    System.out.println("game end! IS DRAW");
                }
            }
        }

    }