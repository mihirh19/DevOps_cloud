# 📜 What is Shell Script?

A shell script is a list of commands in a computer program that is run by the Unix shell 🖥️, which is a command-line interpreter. A shell script usually has comments that describe the steps.

- ✅ Allows you to run if-else statements and loops.
- 📂 A file that contains a series of commands.
- 📝 A plain text file.
- 🚀 It executes the command on each line, one line at a time.
- 🕹️ The terminal usually allows just one command at a time.
- 🧩 Shell script allows you to combine and run multiple commands together.

# 🕒 When Do We Need a Shell Script?

- 📝 You need to enter multiple shell commands and will need to do it again in the future.
- 🛠️ If you know how a piece of work needs to be performed, you can put that knowledge in a shell script.
- 🔁 Shell scripts eliminate repetitive tasks through automation.
- 🛡️ Using a good script can reduce the chance of error.
- ⚡ A shell script can perform a task faster than a human can.
- ⏲️ You have to do a task more than once, but it's something that you rarely do.

# 🌟 Why Use Shell?

- 🚀 Speed of deployment.
- 🔧 No worrying about low-level programming objects.
- 🏃 Ease and speed of learning.
- 🏎️ Performance and efficiency.
- 🤖 Anything you can do on the command line can be automated by writing a shell script.
- 🔄 Can automate tedious or repetitive tasks.
- 🤝 Allow you to hand off work to others.
- 📝 Acts as a form of documentation.
- 🛠️ Fairly quick and easy to write.

# 🔗 Shell Script Shebang

The `#!` syntax is used in scripts to indicate an interpreter for execution under UNIX / Linux operating systems. The directive must be the first line in the Linux shell script and must start with shebang `#!`.

The sharp sign `#` and the bang sign `!` are why it's called the shebang. (sharpbang)

Shebang starts with `#!` characters and the path to the bash or other interpreter of your choice. Make sure the interpreter is the full path to a binary file. For example: `/bin/bash.`

```
#!/bin/bash
sleep 300
```

- ❗ If a script does not contain a shebang, the commands are executed using your shell.
- 🖥️ You can print the `SHELL` variable to show which shell you are using.
- 📝 Different shells have slightly varying syntax.
- 📂 You can see all the available shells in /etc/shells file in Linux operating systems.
- ⚙️ It makes shell scripts more like actual executable files, because they can be the subject of 'exec'.
- 🔍 If you do a `ps` while such a command is running, the real name appears instead of `sh` or `bash`. Likewise, system accounting is done based on the real name.
- 🔄 It will allow other interpreters to fit in more smoothly.

```
#!/usr/bin/python
print("Hello this is python script")
```

```sh
{16:22}/Devops_Cloud/shellscript:main ✗ ➭ chmod +x second.sh

{16:25}/Devops_Cloud/shellscript:main ✗ ➭ ./second.sh
Hello this is python script
```

# 🎨 Print Hello World in Different Colors with `echo` Command

The `echo` command prints all the parameters passed to it on the screen. With the help of `echo` , you can also print text in different colors using ANSI escape codes.

### 📝 Example Script:

```bash
#!/bin/bash

echo 'this is our first                 shellscript'

# Colored messages
echo -e "\033[0;31m❌ fail message"       # Red color
echo -e "\033[0;32m✅ success message"    # Green color
echo -e "\033[0;33m⚠️ warning message here" # Yellow color

```

### 📊 Output Explanation:

- The fail message is printed in red (\033[0;31m), indicating an error or failure.
- The success message is printed in green (\033[0;32m), signifying success.
- The warning message is printed in yellow (\033[0;33m), to highlight caution.
- You'll see the following on your terminal:

- ❌ fail message here in red color.
- ✅ success message here in green color.
- ⚠️ warning message here in yellow color.

This script shows how to make your shell output more informative and visually appealing by using colors. 🌈

# 📝 Comments and Escape Characters.

## 📌 Comments:

Comments provide useful information that helps developers and readers understand the source code. In shell scripting, there are two types of comments:

1. Single-line Comment:

   - A single-line comment starts with a hashtag symbol (#) with no white spaces and lasts till the end of the line.
   - If the comment exceeds one line, place a hashtag on the next line and continue the comment.

   ```bash
   #!/bin/bash
   # This is a single-line comment.
   ```

2. Multi-line Comment:

   - There are some tricky ways to add multi-line comments in shell scripting. Here are a few methods:
   - Method 1: Using `:` and `'`

   ```bash
   #!/bin/bash
   : '
   This is
   a multi-line comment
   in shell scripting.
   '
   ```

   Method 2: Using `:` and `<<'END_COMMENT'`

   ```bash
   #!/bin/bash
   : <<'END_COMMENT'
   This is
   a multi-line comment
   in shell scripting.
   END_COMMENT

   ```

## 🔑 Escape Characters:

An escape character (`\`) tells the shell to interpret the following character literally.

- A non-quoted backslash (`\`) is the escape character. It preserves the literal value of the next character, except for a newline.
- If a backslash is followed by a newline, the newline is treated as a line continuation, allowing you to break long commands across multiple lines for better readability.

## 🧑‍💻 Example Script:

```bash

#!/bin/bash

# purpose: Print some echo commands with examples of comments and escape characters.

echo "this is Mihir Hadavani" # In-line comment
echo 'this is our first                 shellscript' # Another comment

# This is another comment in the shell script.
# The following lines are commented out:
# echo -e "\033[0;31m fail message # here"
# echo -e "\033[0;32m success message # here"
# echo -e "\033[0;33m warning message here"

echo "my
name
is
mihir"

echo "
########### Script Starting ########
purpose: going to install nginx
####################################
"

# Strong quotes (single quotes):
echo 'my \
name \
is \
mihir'   # The escape character \ is taken literally due to strong quoting.

# Newline escaped:
echo " my \
name \
is \
mihir \
"

# Tab, vertical tab, and newline examples:
echo -e "this is mihir \t hadavani \t test name" # Tab
echo -e "this is mihir \v hadavani \v test name" # Vertical tab
echo -e "this is mihir  \n hadavani \n test name" # Newline
```

## 🖥️ Running the Script:

```sh
{16:40}/Devops_Cloud/shellscript:main ✗ ➭ ./escape.sh
this is Mihir Hadavani
this is our first                 shellscript
my
name
is
mihir

########### Script Starting ########
purpose: going to install nginx
####################################

my \
name \
is \
mihir
 my name is mihir
this is mihir    hadavani        test name
this is mihir
               hadavani
                         test name
this is mihir
 hadavani
 test name

```
# 🧠 User-Defined Variables

## 📌 What is a Variable?
A variable is a character string to which you can assign a value. In shell scripting, variables act as pointers to actual data. The shell allows you to create, assign, and delete variables.

## 📋 Rules for Naming Variables:
- Variable names can contain letters (a-z, A-Z), numbers (0-9), or underscores (`_`).
- Reserved words cannot be used as variable names.
- No whitespace is allowed within variable names.
- Special characters like `!`, `*`,`-`, etc., cannot be used.
- The first character of the variable name cannot be a number.
- By convention, shell variable names are often in UPPERCASE.

## 🛠️ Defining Variables:

```bash
MY_MESSAGE="Hello World"
```

🚨 Note: There must be no spaces around the = sign. VAR=value works, but VAR = value does not.

## 📝 Example [variable.sh](./variable.sh)

```bash
#!/bin/bash
# variable.sh
# User-defined variables.

name="Mihir"
age="20"

echo ${name}
echo "My name is ${name} and my age is ${age}"
# echo 'My name is ${name} and my age is ${age}'

work="programm"
var="ing"

# Concatenating variables
echo "I am $working"           # Undefined variable, will result in an empty output
echo "I am ${work}ing"         # Correct usage, will print "programming"
echo "I am ${work}${var}"      # Correct usage, will also print "programming"

```

# 🖥️ System Variables

## 🌐 What are System Variables?
   System variables are created and maintained by the Linux bash shell itself. They are usually defined in CAPITAL LETTERS and are used to configure various aspects of the shell environment, such as`PS1`, `PATH`, `LANG`, `HISTSIZE`, and `DISPLAY`.
   
   ### 🛠️ Commands to Manage Environment Variables:
- `env` : Run a program in a custom environment or print the current environment variables when used without arguments.

- `printenv` :  Print all or the specified environment variables.

- `set` : Set or unset shell variables. It lists all variables, including environment, shell variables, and shell functions.

- `unset` : Delete shell and environment variables.

- `export` : Set environment variables.

### 🌟 Common System Variables:
- `USER`: The current logged-in user.

- `HOME`: The home directory of the current user.
- `EDITOR`: The default file editor used when you type edit in the terminal.
- `SHELL`: The path to the current user's shell, like bash or zsh.
- `LOGNAME`: The name of the current user.
- `PATH`: A list of directories to be searched when executing commands.
- `LANG`: The current locale settings.
- `TERM`: The current terminal emulation.

## 📝 Example [systemvariables.sh](./systemvariables.sh) :
```bash

#!/bin/bash
# 🖥️ System-defined variables.

# Print the home directory of the current user
echo "🏠 Home Directory: ${HOME}"

# Print the type of operating system
echo "🖥️ OS Type: ${OSTYPE}"

# Print the PATH variable
echo "📂 Path: $PATH"

# Print the process ID of the current shell
echo "🔢 Process ID: $$"

# Print the parent process ID
echo "🔙 Parent Process ID: ${PPID}"

# Print the present working directory
echo "📁 Current Directory: $PWD"

# Print the hostname of the machine
echo "💻 Hostname: $HOSTNAME"

# Print the user ID
echo "🆔 User ID: $UID"

# Sleep for 5 seconds
sleep 5

# Print the number of seconds since the script started
echo "⏱️ Seconds Since Start: ${SECONDS}"

```