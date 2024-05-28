# Variables
SRCDIRS := kyngdum
DISTDIR := dist


# Targets
.PHONY: all clean $(SRCDIRS)

all: $(SRCDIRS) 

kyngdum:
	cd $@ && go build -o ../$(DISTDIR)/$@

clean:
	rm -rf $(DISTDIR)/*