import java.util.Collection;
import java.util.ArrayList;

class BadInheritance {
    // START_TASK OMIT
    class Task {
        private String message;

        public Task(String message) {
            this.message = message;
        }

        public void run() {
            System.out.println("running " + this.message);
        }
    }
    // END_TASK OMIT

    // START_RUNNER OMIT
    class Runner {
        private String name;

        public Runner(String name) {
            this.name = name;
        }

        public String getName() {
            return this.name;
        }

        public void run(Task task) { // HL
            task.run();
        }

        public void runAll(Task[] tasks) { // HL
            for (Task task : tasks) {
                run(task);
            }
        }
    }
    // END_RUNNER OMIT

    // START_COUNTING OMIT
    class RunCounter extends Runner {
        private int count;

        public RunCounter(String message) {
            super(message);
            this.count = 0;
        }

        @Override public void run(Task task) {
            count++; // HL
            super.run(task);
        }

        @Override public void runAll(Task[] tasks) {
            count += tasks.length; // HL
            super.runAll(tasks);
        }

        public int getCount() {
            return count;
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
        new BadInheritance().test();
    }
}