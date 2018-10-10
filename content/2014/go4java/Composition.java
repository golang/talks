import java.util.Collection;
import java.util.ArrayList;

class Composition {
    class Task {
        private String message;

        public Task(String message) {
            this.message = message;
        }

        public void run() {
            System.out.println("running " + this.message);
        }
    }

    class Runner {
        private String name;

        public Runner(String name) {
            this.name = name;
        }

        public String getName() {
            return this.name;
        }

        public void run(Task task) {
            task.run();
        }

        public void runAll(Task[] tasks) {
            for (Task task : tasks) {
                run(task);
            }
        }
    }

    // START_COUNTING OMIT
    class RunCounter {
        private Runner runner; // HL
        private int count;

        public RunCounter(String message) {
            this.runner = new Runner(message);
            this.count = 0;
        }

        public void run(Task task) {
            count++;
            runner.run(task);
        }

        public void runAll(Task[] tasks) {
            count += tasks.length;
            runner.runAll(tasks);
        }

        // continued on next slide ...

        // BREAK_COUNTING OMIT
        public int getCount() {
            return count;
        }

        public String getName() {
            return runner.getName();
        }
    }
    // END_COUNTING OMIT

    public void test() {
        // START_MAIN OMIT
        RunCounter runner = new RunCounter("my runner");

        Task[] tasks = { new Task("one"), new Task("two"), new Task("three")};
        runner.runAll(tasks);

        System.out.printf("%s ran %d tasks\n", runner.getName(), runner.getCount());
        // END_MAIN OMIT
    }

    public static void main(String[] args) {
        new Composition().test();
    }
}
