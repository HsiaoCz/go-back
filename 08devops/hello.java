/**
 * hello
 */
public class hello {
    private String username;
    private int age;

    public void setHello(String username, int age) {
        this.username = username;
        this.age = age;
    }

    public String getHelloname() {
        return this.username;
    }

    public int getHelloage() {
        return this.age;
    }

    public void sayHello() {
        System.out.println("Hello");
    }
}