public class Logic {
    static void PrintField(String[][] field) {
        for (int i = 0; i < field.length; i++) {
            for (int j = 0; j < field[i].length; j++) {
                System.out.print(field[i][j] + " ");
            }
            System.out.println();
        }
    }
    static void PrintRules(String[][] rules) {
        for (int i = 0; i < rules.length; i++) {
            for (int j = 0; j < rules[i].length; j++) {
                System.out.print(rules[i][j] + " ");
            }
            System.out.println();
        }
    }
    static void PrintFieldUser(String[][] field,int pos1,int pos2) {
        for (int i = 0; i < field.length; i++) {
            field[Main.pos1_user][Main.pos2_user] = "o";
            for (int j = 0; j < field[i].length; j++) {
                System.out.print(field[i][j] + " ");
            }
            System.out.println();
        }
    }
    static void PrintFieldComputer(String[][] field,int pos1,int pos2) {
        for (int i = 0; i < field.length; i++) {
            field[Main.pos1_computer][Main.pos2_computer] = "x";
            for (int j = 0; j < field[i].length; j++) {
                System.out.print(field[i][j] + " ");
            }
            System.out.println();
        }
    }
    //check to win method
    static void CheckWin(String[][] massive){
        if(massive[0][0].equals("o") && massive[0][1].equals("o") && massive[0][2].equals("o")){
            PrintWinnerO();
        } else if(massive[0][0].equals("x") && massive[0][1].equals("x") && massive[0][2].equals("x")){
            PrintWinnerX();
        }else if(massive[1][0].equals("o") && massive[1][1].equals("o") && massive[1][2].equals("o")){
            PrintWinnerO();
        }else if(massive[1][0].equals("x") && massive[1][1].equals("x") && massive[1][2].equals("x")){
            PrintWinnerX();
        }else if(massive[2][0].equals("o") && massive[2][1].equals("o") && massive[2][2].equals("o")){
            PrintWinnerO();
        }else if(massive[2][0].equals("x") && massive[2][1].equals("x") && massive[2][2].equals("x")){
            PrintWinnerX();
        }else if(massive[0][0].equals("o") && massive[1][1].equals("o") && massive[2][2].equals("o")){
            PrintWinnerO();
        }else if(massive[0][0].equals("x") && massive[1][1].equals("x") && massive[2][2].equals("x")){
            PrintWinnerX();
        }else if(massive[0][2].equals("o") && massive[1][1].equals("o") && massive[2][0].equals("o")){
            PrintWinnerO();
        }else if(massive[0][2].equals("x") && massive[1][1].equals("x") && massive[2][0].equals("x")){
            PrintWinnerX();
        }else if(massive[0][0].equals("o") && massive[1][0].equals("o") && massive[2][0].equals("o")){
            PrintWinnerO();
        }else if(massive[0][0].equals("x") && massive[1][0].equals("x") && massive[2][0].equals("x")){
            PrintWinnerX();
        }else if(massive[0][2].equals("o") && massive[1][2].equals("o") && massive[2][2].equals("o")){
            PrintWinnerO();
        }else if(massive[0][2].equals("x") && massive[1][2].equals("x") && massive[2][2].equals("x")){
            PrintWinnerX();
        }
    }
    static void PrintWinnerO(){
        Main.game = true;
        System.out.println("game end!\n WINNER IS '--O--'");
    }
    static void PrintWinnerX(){
        Main.game = true;
        System.out.println("game end!\n WINNER IS '--x--'");
    }
    static void Step(int number) {
        Main.pos1_user = 0;
        Main.pos2_user = 0;
        switch (number) {
            case 1 -> {
            }
            case 2 -> Main.pos2_user = 1;
            case 3 -> Main.pos2_user = 2;
            case 4 -> Main.pos1_user = 1;
            case 5 -> {
                Main.pos1_user = 1;
                Main.pos2_user = 1;
            }
            case 6 -> {
                Main.pos1_user = 1;
                Main.pos2_user = 2;
            }
            case 7 -> Main.pos1_user = 2;
            case 8 -> {
                Main.pos1_user = 2;
                Main.pos2_user = 1;
            }
            case 9 -> {
                Main.pos1_user = 2;
                Main.pos2_user = 2;
            }
        }
    }
    static void StepComputer(int number) {
        Main.pos1_computer = 0;
        Main.pos2_computer = 0;
        switch (number) {
            case 1 -> {
            }
            case 2 -> Main.pos1_computer = 1;
            case 4 -> Main.pos2_computer = 1;
            case 5 -> {
                Main.pos1_computer = 1;
                Main.pos2_computer = 1;
            }
            case 6 -> {
                Main.pos1_computer = 1;
                Main.pos2_computer = 2;
            }
            case 7 -> Main.pos1_computer = 2;
            case 8 -> {
                Main.pos1_computer = 2;
                Main.pos2_computer = 1;
            }
            case 9 -> {
                Main.pos1_computer = 2;
                Main.pos2_computer = 2;
            }
        }
    }
}
