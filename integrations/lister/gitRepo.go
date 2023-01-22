package gitRepo

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
)

type Repository struct {
	URL string
	*git.Repository
}

func (r *Repository) Clone() error {
	var err error
	r.Repository, err = git.PlainClone(r.URL, false, &git.CloneOptions{
		URL:               r.URL,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	if err != nil {
		return err

	}
	fmt.Println("Cloned repository:", r.URL)
	return nil

}

func (r *Repository) Pull() error {
	worktree, err := r.Repository.Worktree()
	if err != nil {
		return err

	}
	err = worktree.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		return err

	}
	fmt.Println("Pulled repository:", r.URL)
	return nil

}

/*
This script defines a struct Repository that has two fields: URL and Repository. The Repository field is a pointer to a git repository object that is created when the Clone method is called on a Repository struct. The Clone method uses the go-git PlainClone function to clone the repository, and the Pull method uses the go-git Worktree and Pull functions to pull the repository. The main function creates a list of Repository structs, and then iterates through the list, calling the Clone and Pull methods on each struct.
*/

func main() {
	repos := []*Repository{
		&Repository{URL: "https://github.com/user/repo1.git"},
		&Repository{URL: "https://github.com/user/repo2.git"},
		&Repository{URL: "https://github.com/user/repo3.git"},
	}

	for _, repo := range repos {
		if err := repo.Clone(); err != nil {
			log.Fatalln(err)

		}
		if err := repo.Pull(); err != nil {
			log.Fatalln(err)

		}

	}

}
