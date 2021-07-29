# This file covers comments, variables and printing




# comments are ignored in the code, this is for you to make notes about what the code does
# you can comment a line you are on by pressing control + /
# you can comment multiple lines as well if you have multiple lines selected

# this is a single line comment
print("this stuff below wont get commented because its not on the same line")

# you can comment code print("this stuff below wont get commented because its not on the same line")

"""
Another Comment that is multiple lines
    Another line that is ignored by the program

"""

"""
    SHORTCUT:
    CNTRL + /   comments and uncomments the current line you are one, 
                it will also comment all code you have selected, aka you can highlight multiple lines and CNTRL + / will comment it all\
                
"""

print("Hello") # comments can go here as well


"""
    DATA TYPES:
    Python has several types of basic data types
    
    to simplify there are numbers and words
        - a word is called a string or str for short
        - a number is an integer or int, or float 
    
    There are more than this but we dont care. look here if you want the full list:
    https://www.w3schools.com/python/python_datatypes.asp
"""
# this is a word or string
number = "not a number"
# number or integer
number = 2
# number or float
number = 2.5
# boolean
number = False

"""
    VARIABLES:
    typically you store data types in a variable
    variables are just a name like x or number that holds the data
    
    
"""


variable = 2
number = 5

number = variable # you can assign one variable to another
# number is now 2



"""
    ADVICE:
    Give variable names that explain what they are
        x is not a good variable name, age is a good variable name
        imaging writing some code and not looking at it for 5 years, 
        if you come back and cant read what it's doing it's not good, 
        how is someone else suppose to read what you did if you cant?
"""


"""
    PRINTING
    below this window is another window that holds the console where print prints to
    there are multiple tabs below, the one you care about it the Tab labeled Run

    when you run the program any print statements that the program hits will print in the run window
    if you have too much printing out, just comment some of it out
    you can literally print anything!
"""

# print examples
print(2)
print("Hello")
mylist = ["a", "list", "of", "words"] # literally print anything
print(mylist)



im_going_to_print_this = """
Check this out you can print a triple comment, 
but you cant do it with single line prints"""
print(im_going_to_print_this)

# whats cool about this is it will print it out as you see it in triple quotes with the proper tabs and format
cool_print = """                                                                                                    
                                        /%@@@&#//...*/#&@@@#.                                       
                                  /@&.                         ,@@,                                 
                               (@.                                 /@.                              
                             #&                                      .@.                            
                            @                                          #%                           
                           @                                            %#                          
                          /%   @,                                   /@   @                          
                          %/   &                                     @   @,                         
                          .&   @                                    ,@   @                          
                           @   .@                                  .@   ,&                          
                            @   @    ,/%@@@@@@*       #@@@@@&#/.   *%  ,@                           
                             @/ @  @@@@@@@@@@@@(     &@@@@@@@@@@@@ ,& @(                            
                               @@ .@@@@@@@@@@@@      *@@@@@@@@@@@@  @(                              
                  %@@&         @   .@@@@@@@@@#         @@@@@@@@@@   ,%        /@#@/                 
                 #(   *&       @     #@@@@@%   *@@.@@    @@@@@@/    ,&      .@    @                 
                .@*     @@.    ##             &@@@.@@@/             &/    *@(     %@                
              /@            #@& .@#          *@@@@.@@@@          .&@ (@@.           .@              
              (@#**(%&%          /@@/#@@,    .@@@@*@@@@     (@@,&@@.         ,%@%*,*%@              
                         *@@*       @,&( @.               /& @,%(       (@@,                        
                              .@@/   @ @ %@@%(,.    .,/#&@@,#/,@   %@%                              
                                   .&@ *@/&/%,@*% @ @ @ @,@,@ .@/                                   
                                   *&@  @&%#(*@%@%@%@/@ @ @&(  @%.                                  
                              *&@*   @.    #@@@%&*@.@(@@@*    #(   *@@*                             
                   &@&#@@@@%,         @(                     @#         ,&@@@%#@@/                  
                  .@.            *@@,   &&                *@/   *@&,            ,@                  
                     &%      &@*           .(&@@@&&@@@@%*            (@#      @#                    
                      @.   &#                                           &(   .@                     
                      *@,%&                                              .@#,@                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        
"""

print(cool_print)

insert = "This will print"
insert2 = 2021
fString = f"""
    check this out {insert}
    I wrote this in {insert2}
"""

print(fString)

# print f strings are a way to pass functions or variables into strings
test_string = "this will print"

print_this = f"""
    {print(test_string)} and after this will print
"""

print(print_this)





"""
    Printing is probably the most useful tool you have, if you dont know what is happening, print it out! or use the debugger..
"""

