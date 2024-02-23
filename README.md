# # Go Command Line Tools


## Description

[A collection of command lines tools developed in Go programming language!](https://github.com/dx-zone/go-command-line-tools)

This is a delightful assembly of command-line wizardry crafted in the enchanting Go programming language. It's my secret arsenal to bring order to the chaos of the ever-shifting realms of DevOps, making my life more magical and, dare I say, downright entertaining!

## Tools
---
## Log Line Splitter
- **log-line-splitter.go**
Fed up with unraveling the mysteries of failed system logs? Simply toss those logs into a text file, and behold! This nifty tool will magically transform the chaos into a beautifully formatted masterpiece, making decoding errors a breeze. It's like turning "blah blah blah error\nwtf\nexception\nunknown" into a log ballet. 💃✨

The main purpose of the provided Go code, named `log-line-splitter`, is to read the lines of any given ASCII/text file and split the lines based on a specified character or use the default newline character (`\n`). It provides a simple utility for parsing and formatting log files or text documents.

### Functionality

1. **Reading and Splitting Lines:**
   - The program can be executed with one or two command-line arguments.
   - If executed with one argument (`log-line-splitter <filename>`), it reads the content of the specified file, replaces literal `"\n"` with newline characters, and prints each line with a line number.
   - If executed with two arguments (`log-line-splitter <filename> <character-to-split-line>`), it reads the content of the specified file and splits lines based on the specified character.

2. **Command-Line Usage:**
   - If no or more than two arguments are provided, the program displays usage instructions, explaining how to use the tool and providing examples.

3. **ANSI Colors for Output:**
   - The program uses ANSI escape codes for colors (green) in the console output to enhance readability.

4. **Error Handling:**
   - It provides error messages if there are issues reading the file, and it exits with an error code (`os.Exit(1)`) in case of errors.

5. **Additional Information:**
   - In the case of incorrect command-line arguments, the program outputs information about its purpose and usage.

### Usage Examples

- Execute with one argument:
  ```bash
  log-line-splitter my.txt
  log-line-splitter my.txt "\t"
  log-line-splitter my.txt "c"
  log-line-splitter my.txt "xyz"

<!-- - TODO: Create a recorded GIF image showcasing the use of this app -->
---

## DNS Zone Helper
- **dns-zone-helper.go**

Exhausted from the mental gymnastics of remembering DNS commands and playing hide-and-seek with those elusive config files? Fear not! This tool not only conjures up the zone config files like a DNS wizard but also acts as your personal memory jogger for those command acrobatics needed to wrangle those DNS zones into submission. It's DNS assistance with a sprinkle of magic and a dash of humor! 🧙‍♂️✨

The main purpose of the tool is to provide information and commands related to a DNS zone files in a BIND server. It is a command-line utility that takes input parameters, including the domain name and view name, and provides various details and commands associated with that DNS zone.

The tool `dns-zone-helper`, serves as a utility for managing and retrieving information related to DNS (Domain Name System) zone files. It provides commands to inspect, configure, and perform various actions on DNS zones within a specified BIND (Berkeley Internet Name Domain) server environment.

## Functionality

1. **Zone Information Display:**
   - The primary function is to display information about a DNS zone, including its name, directory, view, filename, and absolute path.

2. **Command-Line Interface:**
   - The program expects at least two command-line arguments: `zone-name` and `view-name`.
   - It provides a usage message if the required arguments are not provided.

3. **Options:**
   - The tool supports additional options based on the provided command-line arguments:
     - Displaying suggested commands to search for DNS zone configuration (`config` option).
     - Displaying a set of suggested commands for working with the DNS zone (`commands` option).

4. **Usage Examples:**
   - Examples of valid command-line arguments and their purposes are provided in the usage message.

## Commands and Actions

The program generates and executes commands related to the specified DNS zone, including:
- Checking BIND configuration (`named-checkconf`).
- Searching for configuration options (`grep` command).
- Verifying the existence of the zone file (`sudo ls` command).
- Checking the syntax of the zone configuration (`named-checkconf` and `named-checkzone`).
- Performing actions like freezing, thawing, and reloading the DNS zone (`rndc` commands).

## Usage Examples

- Display information about the DNS zone:
  ```bash
  dns-zone-helper example.net external
  dns-zone-helper example.net my-view

<!-- - TODO: Create a recorded GIF image showcasing the use of this app -->
---
## Password Complexity Check
- **password-complexity-check**

Ever feel like your passwords are staging a rebellion against complexity standards? Enter this superhero application, swooping in to validate if your chosen password passes the "according to me" complexity criteria. Say goodbye to frustration and hello to password peace – because sometimes, it's good to be the superhero of your own security saga! 🦸‍♂️🔐

The main purpose of the **password-complexity-check** is to determine whether a given password meets certain complexity criteria. The code prompts the user to enter a password, reads the input, and then evaluates whether the password satisfies the following conditions:

1. **Minimum Length:**
     - The password must be at least 8 characters long.
2. **Character Requirements:**

     - The password must contain at least one uppercase letter.
     - The password must contain at least one lowercase letter.
     - The password must contain at least one digit.
     - The password must contain at least one special character (non-alphanumeric).
3. **Output:**
     - If the entered password satisfies all the specified complexity criteria, the program prints: "Password is complex and meets the criteria."
     - If the entered password does not meet the complexity criteria or is less than 8 characters, the program prints: "Password does not meet the complexity criteria or is less than 8 characters."
This code is a simple password complexity checker, providing a basic level of security validation for user-generated passwords. It checks for a combination of factors to ensure that passwords are sufficiently strong and meet standard security guidelines.

<!-- - TODO: Create a recorded GIF image showcasing the use of this app -->

---
<!-- TODO: Include the Go tool to find publics and private IP addresses in text/DNS record files. -->
<!-- - TODO: Include the title of the application here -->
<!-- - TODO: Create a recorded GIF image showcasing the use of this app -->


### Requirements
- Ready to dive into the world of Go? Go ahead and install Go (pun totally intended!). Go is the superhero programming language behind these source codes. You can either make them flex their muscles instantly with ```go run source_code.go``` (no need to compile) or give them a power boost with ```go build source_code.go``` to create a binary sidekick in your current working directory. Let the coding adventures begin!


### Examples

Grab your popcorn and don your patience hat – more documentation is on its way! Stay tuned for the thrilling sequel of additional documentation, coming soon to your screen! 🍿📜
