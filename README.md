# Generics in Go
This is the repository for the LinkedIn Learning course Generics in Go. The full course is available from [LinkedIn Learning][lil-course-url].

![Generics in Go][lil-thumbnail-url] 

Repeating code by writing boilerplate code over and over again can often lead to mistake-ridden, messy code. In this course, Anna-Katharina Wickert shows you how to use generics in Go to write cleaner and easier-to-read code. Anna-Katharina shows how you can use generics—long on developers’ wishlists and introduced in Go 1.18—to write functions and types that work with a set of types instead of one concrete type. She covers type parameters, typesets, and type inference, and puts it all into practice with examples and looks into new packages shipped along with generics.

## Instructions
This repository has branches for each of the videos in the course. You can use the branch pop up menu in github to switch to a specific branch and take a look at the course at that stage, or you can add `/tree/BRANCH_NAME` to the URL to go to the branch you want to access.

## Branches
The branches are structured to correspond to the videos in the course. The naming convention is `CHAPTER#_MOVIE#`. As an example, the branch named `02_03` corresponds to the second chapter and the third video in that chapter. 
Some branches will have a beginning and an end state. These are marked with the letters `b` for "beginning" and `e` for "end". The `b` branch contains the code as it is at the beginning of the movie. The `e` branch contains the code as it is at the end of the movie. The `main` branch holds the final state of the code when in the course.

When switching from one exercise files branch to the next after making changes to the files, you may get a message like this:

    error: Your local changes to the following files would be overwritten by checkout:        [files]
    Please commit your changes or stash them before you switch branches.
    Aborting

To resolve this issue:
	
    Add changes to git using this command: git add .
	Commit changes using this command: git commit -m "some message"


## Installing
1. To use these exercise files, you must have the following installed:
	- Go 1.18 or later - [Download and install](https://go.dev/doc/install) 
2. Clone this repository into your local machine using the terminal (Mac), CMD (Windows), or a GUI tool like SourceTree.
3. To install all dependencies run `go get .` from this directory, the root directory. 
Be aware that the dependencies differ between the branches, so ensure to execute the command when you switch to another branch.
4. Run the example program, e.g., with `go run main.go`.


### Instructor

Anna-Katharina Wickert 
                            


                            

Check out my other courses on [LinkedIn Learning](https://www.linkedin.com/learning/instructors/anna-katharina-wickert).

[lil-course-url]: https://www.linkedin.com/learning/generics-in-go?dApp=59033956
[lil-thumbnail-url]: https://media.licdn.com/dms/image/C4E0DAQGQ2F4Fuv3azA/learning-public-crop_675_1200/0/1670989462878?e=2147483647&v=beta&t=TDb41Ax46yS4RGBnyd-349wmZjcqNDUpa9g7Pz7HJW8



