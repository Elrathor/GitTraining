![Git Logo](https://git-scm.com/images/logos/downloads/Git-Logo-2Color.png)

# Introduction to git
Some parts of the following information are targeted to a windows user, although most of it can be applied to any other operating system in a modified form.

## 0 - About Git
> By far, the most widely used modern version control system in the world today is Git. Git is a mature, actively maintained open source project originally developed in 2005 by Linus Torvalds, the famous creator of the Linux operating system kernel. A staggering number of software projects rely on Git for version control, including commercial projects as well as open source. Developers who have worked with Git are well represented in the pool of available software development talent and it works well on a wide range of operating systems and IDEs (Integrated Development Environments). 

> Having a distributed architecture, Git is an example of a DVCS (hence Distributed Version Control System). Rather than have only one single place for the full version history of the software as is common in once-popular version control systems like CVS or Subversion (also known as SVN), in Git, every developer's working copy of the code is also a repository that can contain the full history of all changes.

Source: https://www.atlassian.com/git/tutorials/what-is-git

## 1 - Tools & Requirements For Using Git
- Git for Windows: https://git-scm.com/download/win
- (optional - recomended) - Graphical git client:  https://desktop.github.com/
- (optional - recomended) - Editor with git capabilities: https://code.visualstudio.com/download 

## 2 - Ressources
These ressources can be used as further reading material or command reference
- MIT online lecture: https://www.youtube.com/watch?v=2sjqTHE0zok
- Git reference book: https://git-scm.com/book/en/v2 
- Git reference book (de): https://git-scm.com/book/de/v2
- Basic tutorial: https://www.freecodecamp.org/news/what-is-git-and-how-to-use-it-c341b049ae61
- Cheat Sheet: http://cheat.sh/git

## 3 - Super Basic Flow
1. Create a repository
   ```BASH
   git init
   ```
2. Create and/or modify a text file (eg. programm or even a thesis)
3. Add changes to the list of staged changes
   ```BASH
   git add .
   ```
4. Commit chages to the change tree
   ```BASH
   git commit -m "<INSERT COMMENT ABOUT WHAT CHANGED>"
   ```
5. Go to step 2 and repeat
