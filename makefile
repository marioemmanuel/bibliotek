# PROJECT: bibliotek makefile
# AUTHOR: MARIO EMMANUEL
# DATE: 2023/DEC/19
# github.com/marioemmanuel/bibliotek

bibliotek:	bibliotek.go filetree.go render.go www.go 
			go build bibliotek.go filetree.go render.go www.go

clean:		;
			rm bibliotek
