import random
import unittest
from typing import Tuple

"""
Reference: https://docs.python.org/3.2/library/unittest.html
Basics:
unittest supports test automation, sharing of setup and shutdown code for tests, aggregation of tests into collections, 
and independence of the tests from the reporting framework. The unittest module provides classes that make it easy to 
support these qualities for a set of tests.

To achieve this, unittest supports some important concepts:

test fixture
A test fixture represents the preparation needed to perform one or more tests, and any associate cleanup actions. 
This may involve, for example, creating temporary or proxy databases, directories, or starting a server process.
test case
A test case is the smallest unit of testing. It checks for a specific response to a particular set of inputs. unittest 
provides a base class, TestCase, which may be used to create new test cases.
test suite
A test suite is a collection of test cases, test suites, or both. It is used to aggregate tests that should be executed 
together.

test runner
A test runner is a component which orchestrates the execution of tests and provides the outcome to the user. The runner 
may use a graphical interface, a textual interface, or return a special value to indicate the results of executing the tests.
The test case and test fixture concepts are supported through the TestCase and FunctionTestCase classes; the former 
should be used when creating new tests, and the latter can be used when integrating existing test code with a 
unittest-driven framework. When building test fixtures using TestCase, the setUp() and tearDown() methods can be 
overridden to provide initialization and cleanup for the fixture. With FunctionTestCase, existing functions can be 
passed to the constructor for these purposes. When the test is run, the fixture initialization is run first; if it 
succeeds, the cleanup method is run after the test has been executed, regardless of the outcome of the test. Each 
instance of the TestCase will only be used to run a single test method, so a new fixture is created for each test.

Test suites are implemented by the TestSuite class. This class allows individual tests and test suites to be aggregated; 
when the suite is executed, all tests added directly to the suite and in “child” test suites are run.

A test runner is an object that provides a single method, run(), which accepts a TestCase or TestSuite object as a 
parameter, and returns a result object. The class TestResult is provided for use as the result object. unittest provides 
the TextTestRunner as an example test runner which reports test results on the standard error stream by default. 
Alternate runners can be implemented for other environments (such as graphical environments) without any need to 
derive from a specific class.

"""


class TestSequenceFunctions(unittest.TestCase):
    """
    A testcase is created by subclassing unittest.TestCase.
    The three individual tests are defined with methods whose names start with the letters test.
    This naming convention informs the test runner about which methods represent tests.

    The crux of each test is a call to assertEqual() to check for an expected result; assertTrue() to verify a condition;
    or assertRaises() to verify that an expected exception gets raised. These methods are used instead of the assert
    statement so the test runner can accumulate all test results and produce a report.

    When a setUp() method is defined, the test runner will run that method prior to each test. Likewise, if a tearDown()
    method is defined, the test runner will invoke that method after each test. In the example, setUp() was used to
    create a fresh sequence for each test.
    """

    def setUp(self):
        """Runs multiple times for each function that calls it"""
        self.seq = list(range(10))
        print(f"functon: setUp(), runs for each test function"
              f"\n  Seq: {self.seq}")

    def test_shuffle(self):
        # make sure the shuffled sequence does not lose any elements
        random.shuffle(self.seq)
        print(f"Shuffled Values: {self.seq}")
        self.seq.sort()
        print(f"Resorted Values: {self.seq}")
        # checks if two list are equal
        self.assertEqual(self.seq, list(range(10)))

        # should raise an exception for an immutable sequence
        self.assertRaises(TypeError, random.shuffle, (1, 2, 3))

    def test_choice(self):
        """check if choice in list(0-9) created in self.seq during setUp"""
        element = random.choice(self.seq)
        # intentional fail
        # element = 12
        print(f"function: test_choice random selection: {element}")

        # Checks if element value is in seq
        self.assertTrue(element in self.seq)

    def test_sample(self):
        with self.assertRaises(ValueError):
            random.sample(self.seq, 20)
        for element in random.sample(self.seq, 5):
            self.assertTrue(element in self.seq)

"""
simple way to run the tests. unittest.main() provides a command-line interface to the test script. 
When run from the command line, the above script produces an output that looks like this:
"""
if __name__ == '__main__':
    unittest.main()

"""
Testing from CLI
The unittest module can be used from the command line to run tests from modules, classes or even individual test methods:

python -m unittest test_module1 test_module2
python -m unittest test_module.TestClass
python -m unittest test_module.TestClass.test_method

You can pass in a list with any combination of module names, and fully qualified class or method names.

Test modules can be specified by file path as well:

python -m unittest tests/test_something.py
This allows you to use the shell filename completion to specify the test module. The file specified must still be 
importable as a module. The path is converted to a module name by removing the ‘.py’ and converting path separators 
into ‘.’. If you want to execute a test file that isn’t importable as a module you should execute the 
file directly instead.

You can run tests with more detail (higher verbosity) by passing in the -v flag:

python -m unittest -v test_module
When executed without arguments Test Discovery is started:

python -m unittest
For a list of all the command-line options:

python -m unittest -h
"""


"""
Test Discovery
Unittest supports simple test discovery. In order to be compatible with test discovery, all of the test files must be 
modules or packages importable from the top-level directory of the project (this means that their filenames must be 
valid identifiers).

CLI
cd project_directory
python -m unittest discover
    note: this is equal to python -m unittest
        If you want to pass arguments to test discovery the discover sub-command must be used explicitly.
"""



class Widget():
    def __init__(self, name, size=(50, 50)):
        self.size = size
        self.name = name
        # print(self.size == (50, 50))

    def resize(self, tup: Tuple):
        self.size = len(tup)



# To make your own test cases you must write subclasses of TestCase or use FunctionTestCase.
# class DefaultWidgetSizeTestCase(unittest.TestCase):
#     def runTest(self):
#         widget = Widget('The widget')
#         print(f"size: {widget.size}")
#         self.assertEqual(widget.size(), (50, 50), 'incorrect default size')



class SimpleWidgetTestCase(unittest.TestCase):
    def setUp(self):
        self.widget = Widget('The widget')
        print(f"size: {self.widget.size}")
    """
    If the setUp() method raises an exception while the test is running, the framework will consider the test to have 
    suffered an error, and the runTest() method will not be executed.
    
    Similarly, we can provide a tearDown() method that tidies up after the runTest() method has been run:
    """
    def tearDown(self):
        self.widget.dispose()
        self.widget = None

class DefaultWidgetSizeTestCase(SimpleWidgetTestCase):
    def runTest(self):
        self.assertEqual(self.widget.size(), (50,50),
                         'incorrect default size')

class WidgetResizeTestCase(SimpleWidgetTestCase):
    def runTest(self):
        self.widget.resize(100, 150)
        self.assertEqual(self.widget.size(), (100,150),
                         'wrong size after resize')
