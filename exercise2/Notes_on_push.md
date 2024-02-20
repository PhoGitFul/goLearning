mal@ip-10-192-241-136 goLearning % git pull
Already up to date.
mal@ip-10-192-241-136 goLearning % git checkout -b 03_Constant
fatal: a branch named '03_Constant' already exists
mal@ip-10-192-241-136 goLearning % git branch --set-upstream-to=origin/main 03_Constant
branch '03_Constant' set up to track 'origin/main'.
mal@ip-10-192-241-136 goLearning % cd 03_Constants
mal@ip-10-192-241-136 03_Constants % go run main.go        
main.go:1:1: expected 'package', found 'EOF'
mal@ip-10-192-241-136 03_Constants % go run main.go
0
5.888889e+13
float32complex128
float64
float64
225 222
-32767 32765
mal@ip-10-192-241-136 03_Constants % go run main.go
0
5.888889e+13
float32complex128
float64
float64
225 222
32767 -32768 -32767 32765
mal@ip-10-192-241-136 03_Constants % go run main.go
0
5.888889e+13
float32complex128
float64
float64
225 222
-32767 -32766 -32765 32767
mal@ip-10-192-241-136 03_Constants % go run main.go
0
5.888889e+13
float32complex128
float64
float64
225 222
-32767 32767 0 2 -32768 -32767 32765
mal@ip-10-192-241-136 03_Constants % go run main.go
# command-line-arguments
./main.go:30:23: syntax error: unexpected negY in argument list; possibly missing comma or )
mal@ip-10-192-241-136 03_Constants % go run main.go
0
5.888889e+13
float32complex128
float64
float64
225 222
negY = -32767 Y = 32767
Y+negY = 0 negY-Y = 2 -32768 -32767 32765
Y+1 = -32768 Y+2 = -32767 Y+2 = -32767
mal@ip-10-192-241-136 03_Constants % go run main.go
0
5.888889e+13
float32complex128
float64
float64
225 222
negY = -32767 Y = 32767
Y+negY = 0 negY-Y = 2
Y+1 = -32768 Y+2 = -32767 negY+1 = -32766
mal@ip-10-192-241-136 03_Constants % go run main.go
0
5.888889e+13
float32complex128
float64
float64
225 222
negY = -32767 Y = 32767
Y+negY = 0 negY-Y = 2
Y+0 = 32767 Y+1 = -32768 Y+2 = -32767 negY+1 = -32766 negY-1 = -32768
mal@ip-10-192-241-136 03_Constants % go run main.go
0
5.888889e+13
float32complex128
float64
float64
225 222
negY = -32767 Y = 32767
Y+negY = 0 negY-Y = 2
Y+0 = 32767 Y+1 = -32768 Y+2 = -32767 Y-1 = 32766 Y-2 = 32765
negY+0 = -32767 negY+1 = -32766 negY+2 = -32765 negY-1 = -32768 negY-2 = 32767 -32768
mal@ip-10-192-241-136 03_Constants % go run main.go
0
5.888889e+13
float32complex128
float64
float64
225 222
negY = -32767 , Y = 32767
Y+negY = 0 , negY-Y = 2
Y+0 = 32767 , Y+1 = -32768 , Y+2 = -32767 , Y-1 = 32766 , Y-2 = 32765
negY+0 = -32767 , negY+1 = -32766 , negY+2 = -32765 , negY-1 = -32768 negY-2 = 32767 -32768
mal@ip-10-192-241-136 03_Constants % go run main.go
0
5.888889e+13
float32complex128
float64
float64
225 222
negY = -32767 , Y = 32767
Y+negY = 0 , negY-Y = 2 -32767 - 32767 => -32768 - 32766
sort of flips to positive as can't get any lower than -32768 so is like +32768 - 32766 = 2
Y+0 = 32767 , Y+1 = -32768 , Y+2 = -32767 , Y-1 = 32766 , Y-2 = 32765
negY+0 = -32767 , negY+1 = -32766 , negY+2 = -32765 , negY-1 = -32768 negY-2 = 32767 -32768
mal@ip-10-192-241-136 03_Constants % go run main.go
# command-line-arguments
./main.go:31:58: syntax error: unexpected comma, expected expression
mal@ip-10-192-241-136 03_Constants % go run main.go
0
5.888889e+13
float32complex128
float64
float64
225 222
negY = -32767 , Y = 32767
Y+negY = 0 , negY-Y = 2 , -32767 - 32767 => -32768 - 32766
sort of flips to positive as can't get any lower than -32768 so is like +32768 - 32766 = 2
Y+0 = 32767 , Y+1 = -32768 , Y+2 = -32767 , Y-1 = 32766 , Y-2 = 32765
negY+0 = -32767 , negY+1 = -32766 , negY+2 = -32765 , negY-1 = -32768 negY-2 = 32767 -32768
mal@ip-10-192-241-136 03_Constants % go run main.go
e = 0
5.888889e+13
float32complex128
float64
float64
225 222
negY = -32767 , Y = 32767
Y+negY = 0 , negY-Y = 2 , -32767 - 32767 => -32768 - 32766
sort of flips to positive as can't get any lower than -32768 so is like +32768 - 32766 = 2
Y+0 = 32767 , Y+1 = -32768 , Y+2 = -32767 , Y-1 = 32766 , Y-2 = 32765
negY+0 = -32767 , negY+1 = -32766 , negY+2 = -32765 , negY-1 = -32768 negY-2 = 32767 -32768
mal@ip-10-192-241-136 03_Constants % go run main.go
e = 0
f = 5.888889e+13
f Type =
float32g Type = complex128
h Type = float64
i Type = float64
225 222
negY = -32767 , Y = 32767
Y+negY = 0 , negY-Y = 2 , -32767 - 32767 => -32768 - 32766
sort of flips to positive as can't get any lower than -32768 so is like +32768 - 32766 = 2
Y+0 = 32767 , Y+1 = -32768 , Y+2 = -32767 , Y-1 = 32766 , Y-2 = 32765
negY+0 = -32767 , negY+1 = -32766 , negY+2 = -32765 , negY-1 = -32768 negY-2 = 32767 -32768
mal@ip-10-192-241-136 03_Constants % go run main.go
e = 0
f = 5.888889e+13
f Type =
float32
g Type = complex128
h Type = float64
i Type = float64
225 222
negY = -32767 , Y = 32767
Y+negY = 0 , negY-Y = 2 , -32767 - 32767 => -32768 - 32766
sort of flips to positive as can't get any lower than -32768 so is like +32768 - 32766 = 2
Y+0 = 32767 , Y+1 = -32768 , Y+2 = -32767 , Y-1 = 32766 , Y-2 = 32765
negY+0 = -32767 , negY+1 = -32766 , negY+2 = -32765 , negY-1 = -32768 negY-2 = 32767 -32768
mal@ip-10-192-241-136 03_Constants % go run main.go
e = 0
f = 5.888889e+13
f Type =
float32
g Type = complex128
h Type = float64
i Type = float64
j Type = complex64
225 222
negY = -32767 , Y = 32767
Y+negY = 0 , negY-Y = 2 , -32767 - 32767 => -32768 - 32766
sort of flips to positive as can't get any lower than -32768 so is like +32768 - 32766 = 2
Y+0 = 32767 , Y+1 = -32768 , Y+2 = -32767 , Y-1 = 32766 , Y-2 = 32765
negY+0 = -32767 , negY+1 = -32766 , negY+2 = -32765 , negY-1 = -32768 negY-2 = 32767 -32768
mal@ip-10-192-241-136 03_Constants % go build main.go
mal@ip-10-192-241-136 03_Constants % ./main
e = 0
f = 5.888889e+13
f Type =
float32
g Type = complex128
h Type = float64
i Type = float64
j Type = complex64
225 222
negY = -32767 , Y = 32767
Y+negY = 0 , negY-Y = 2 , -32767 - 32767 => -32768 - 32766
sort of flips to positive as can't get any lower than -32768 so is like +32768 - 32766 = 2
Y+0 = 32767 , Y+1 = -32768 , Y+2 = -32767 , Y-1 = 32766 , Y-2 = 32765
negY+0 = -32767 , negY+1 = -32766 , negY+2 = -32765 , negY-1 = -32768 negY-2 = 32767 -32768
mal@ip-10-192-241-136 03_Constants % ls
main    main.go
mal@ip-10-192-241-136 03_Constants % ./02_Constants
zsh: no such file or directory: ./02_Constants
mal@ip-10-192-241-136 03_Constants % cd ..
mal@ip-10-192-241-136 goLearning % ls            
01_helloworld                   exercise
02_Variables                    go.mod
03_Constants                    goLearning.code-workspace
README.md
mal@ip-10-192-241-136 goLearning % ./02_Constants
zsh: no such file or directory: ./02_Constants
mal@ip-10-192-241-136 goLearning % cd 02_Constants
cd: no such file or directory: 02_Constants
mal@ip-10-192-241-136 goLearning % cd 03_Constants
mal@ip-10-192-241-136 03_Constants % git init
Initialized empty Git repository in /Users/mal/go/src/github.com/automation-1/projects/goLearning/03_Constants/.git/
mal@ip-10-192-241-136 03_Constants % git status
On branch main

No commits yet

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        main
        main.go

nothing added to commit but untracked files present (use "git add" to track)
mal@ip-10-192-241-136 03_Constants % git add .
mal@ip-10-192-241-136 03_Constants % git commit -m "first commit"
[main (root-commit) 176ea32] first commit
 2 files changed, 42 insertions(+)
 create mode 100755 main
 create mode 100644 main.go
mal@ip-10-192-241-136 03_Constants % git branch -M main
mal@ip-10-192-241-136 03_Constants % git remote add git@github.com:PhoGitFul/goLearning.git
usage: git remote add [<options>] <name> <url>

    -f, --fetch           fetch the remote branches
    --tags                import all tags and associated objects when fetching
                          or do not fetch any tag at all (--no-tags)
    -t, --track <branch>  branch(es) to track
    -m, --master <branch> master branch
    --mirror[=(push|fetch)]
                          set up remote as a mirror to push to or fetch from

mal@ip-10-192-241-136 03_Constants % git push -u origin main
fatal: 'origin' does not appear to be a git repository
fatal: Could not read from remote repository.

Please make sure you have the correct access rights
and the repository exists.
mal@ip-10-192-241-136 03_Constants % git remote add origin git@github.com:PhoGitFul/goLearning.git
mal@ip-10-192-241-136 03_Constants % git push -u origin main
To github.com:PhoGitFul/goLearning.git
 ! [rejected]        main -> main (fetch first)
error: failed to push some refs to 'github.com:PhoGitFul/goLearning.git'
hint: Updates were rejected because the remote contains work that you do not
hint: have locally. This is usually caused by another repository pushing to
hint: the same ref. If you want to integrate the remote changes, use
hint: 'git pull' before pushing again.
hint: See the 'Note about fast-forwards' in 'git push --help' for details.
mal@ip-10-192-241-136 03_Constants % cd ..
mal@ip-10-192-241-136 goLearning % cd exercise2
mal@ip-10-192-241-136 exercise2 % go run exeConstant.go      
Constant Exercise
Gopher Hotel 24.806078 -78.243027 12
mal@ip-10-192-241-136 exercise2 % go run exeConstant.go
Constant Exercise
Hotel: Gopher Hotel is located at longitude - 24.806078 , latitude - -78.243027 , with occupancy of  12
mal@ip-10-192-241-136 exercise2 % go run exeConstant.go
Gopher Hotel 24.806078 -78.243027 12
mal@ip-10-192-241-136 exercise2 % go run exeConstant.go
Gopher Hotel 24.806078 -78.243027 12
float64
float64
mal@ip-10-192-241-136 exercise2 % git status
On branch 03_Constant
Your branch is up to date with 'origin/main'.

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        ../03_Constants/
        ./

nothing added to commit but untracked files present (use "git add" to track)
mal@ip-10-192-241-136 exercise2 % cd ..
mal@ip-10-192-241-136 goLearning % git status
On branch 03_Constant
Your branch is up to date with 'origin/main'.

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        03_Constants/
        exercise2/

nothing added to commit but untracked files present (use "git add" to track)
mal@ip-10-192-241-136 goLearning % git add .
warning: adding embedded git repository: 03_Constants
hint: You've added another git repository inside your current repository.
hint: Clones of the outer repository will not contain the contents of
hint: the embedded repository and will not know how to obtain it.
hint: If you meant to add a submodule, use:
hint: 
hint:   git submodule add <url> 03_Constants
hint: 
hint: If you added this path by mistake, you can remove it from the
hint: index with:
hint: 
hint:   git rm --cached 03_Constants
hint: 
hint: See "git help submodule" for more information.
mal@ip-10-192-241-136 goLearning % git branch --set-upstream-to=origin/main 03_Constant
branch '03_Constant' set up to track 'origin/main'.
mal@ip-10-192-241-136 goLearning % cd 03_Constant
cd: no such file or directory: 03_Constant
mal@ip-10-192-241-136 goLearning % cd 03_Constants
mal@ip-10-192-241-136 03_Constants % git branch --set-upstream-to=origin/main 03_Constant
fatal: branch '03_Constant' does not exist
mal@ip-10-192-241-136 03_Constants % cd ..
mal@ip-10-192-241-136 goLearning % git status
On branch 03_Constant
Your branch is up to date with 'origin/main'.

Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        new file:   03_Constants
        new file:   exercise2/README.md
        new file:   exercise2/exeConstant.go

mal@ip-10-192-241-136 goLearning % git add.
git: 'add.' is not a git command. See 'git --help'.

The most similar command is
        add
mal@ip-10-192-241-136 goLearning % git add .
mal@ip-10-192-241-136 goLearning % git status
On branch 03_Constant
Your branch is up to date with 'origin/main'.

Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        new file:   03_Constants
        new file:   exercise2/README.md
        new file:   exercise2/exeConstant.go

mal@ip-10-192-241-136 goLearning % git commit -m "03_Constants and Exercise created"
[03_Constant c093206] 03_Constants and Exercise created
 3 files changed, 36 insertions(+)
 create mode 160000 03_Constants
 create mode 100644 exercise2/README.md
 create mode 100644 exercise2/exeConstant.go
mal@ip-10-192-241-136 goLearning % git push origin
fatal: The upstream branch of your current branch does not match
the name of your current branch.  To push to the upstream branch
on the remote, use

    git push origin HEAD:main

To push to the branch of the same name on the remote, use

    git push origin HEAD

To choose either option permanently, see push.default in 'git help config'.

To avoid automatically configuring an upstream branch when its name
won't match the local branch, see option 'simple' of branch.autoSetupMerge
in 'git help config'.

mal@ip-10-192-241-136 goLearning % git push origin HEAD
Enumerating objects: 6, done.
Counting objects: 100% (6/6), done.
Delta compression using up to 10 threads
Compressing objects: 100% (5/5), done.
Writing objects: 100% (5/5), 1.28 KiB | 1.28 MiB/s, done.
Total 5 (delta 1), reused 0 (delta 0), pack-reused 0
remote: Resolving deltas: 100% (1/1), completed with 1 local object.
remote: 
remote: Create a pull request for '03_Constant' on GitHub by visiting:
remote:      https://github.com/PhoGitFul/goLearning/pull/new/03_Constant
remote: 
To github.com:PhoGitFul/goLearning.git
 * [new branch]      HEAD -> 03_Constant
mal@ip-10-192-241-136 goLearning % 