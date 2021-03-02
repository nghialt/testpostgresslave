package read;

import java.sql.*;

import static util.Constant.SLAVE_DSN;

public class Main {
    public static void main(String[] args) {
        Connection c = null;
        Statement stmt = null;
        int j = 0;
        try {
            Class.forName("org.postgresql.Driver");
            c = DriverManager.getConnection(SLAVE_DSN);

            c.setAutoCommit(false);
            System.out.println("Opened database successfully");

            for (j = 0; j < 20; j++ ){
                stmt = c.createStatement();
                ResultSet rs = stmt.executeQuery("SELECT * FROM CUSTOMERS;");
//                ResultSet rs = stmt.executeQuery("SELECT * FROM CUSTOMERS WHERE ID < 20000;");
                long step = 1;
                while (step < 5) {
                    step = System.currentTimeMillis() % 10;
                    Thread.sleep(145);
                }
                int i = 0;
                while (rs.next()) {
                    i++;
                    if (i % step != 0) {
                        continue;
                    }
                    int id = rs.getInt("id");
                    String name = rs.getString("first_name");
                    int count = rs.getInt("count");
                    System.out.println("ID = " + id);
                    System.out.println("NAME = " + name);
                    System.out.println("COUNT = " + count);
                    System.out.println();
                }
                rs.close();
                stmt.close();
                System.out.println("step " + step);
            }
        } catch (SQLException | ClassNotFoundException | InterruptedException throwables) {
            throwables.printStackTrace();
            System.out.println("stop at " + j);
        } finally {
            if (c != null) {
                try {
                    c.close();
                } catch (SQLException throwables) {
                    throwables.printStackTrace();
                }
            }
        }
    }
}
