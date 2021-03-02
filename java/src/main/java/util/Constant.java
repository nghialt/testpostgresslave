package util;

public interface Constant {
    String MASTER_DSN = "jdbc:postgresql://localhost:55004/database?user=postgres&password=password&ssl=false";
    String SLAVE_DSN = "jdbc:postgresql://localhost:55004/database?user=postgres&password=password&ssl=false";
}
