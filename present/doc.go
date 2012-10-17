// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Present displays slide presentations.
It runs a web server that presents slide files from the current directory.

It may be run as a stand-alone command or an App Engine app.
The stand-alone version permits the execution of programs from within a slide
presentation. The App Engine version does not provide this functionality.

Usage of present:
  -base="": base path for slide template and static resources
  -http="127.0.0.1:3999": host:port to listen on
  -template="": alternate slide template file

You may use the app.yaml file provided in the root of the go.talks repository
to deploy present to App Engine:
	appcfg.py update -A your-app-id -V your-app-version /path/to/go.talks

The file slide format

Slide files have the following format.  The first non-blank non-comment
line is the title, so the header looks like

	Title of presentation
	Subtitle of presentation
	<blank line>
	Presenter Name
	Job title, Company
	joe@example.com
	http://url/
	@twitter_name

The presenter section may contain a mixture of text, twitter names, and links.
Only the plain text lines will be displayed on the presentation front page.

Multiple presenters may be specified, separated by a blank line.

After that come slides/sections, each after a blank line:

	* Title of slide or section (must have asterisk)

	Some Text

	- bullets
	- more bullets
	- a bullet with

	Some More text

	  Preformatted text
	  is indented (however you like)

	Further Text, including invocations like:

	.code x.go /^func main/,/^}/
	.play y.go
	.image image.jpg
	.link http://foo label
	.html file.html

	Again, more text

Blank lines are OK (not mandatory) after the title and after the
text.  Text, bullets, and .code etc. are all optional; title is
not.

Lines starting with # in column 1 are commentary.

Fonts:

Within the input for plain text or lists, text bracketed by font
markers will be presented in italic, bold, or program font.
Marker characters are _ (italic), * (bold) and ` (program font).
Unmatched markers appear as plain text.
Within marked text, a single marker character becomes a space
and a doubled single marker quotes the marker character.

	_italic_
	*bold*
	`program`
	_this_is_all_italic_
	_Why_use_scoped__ptr_? Use plain ***ptr* instead.

Functions:

A number of template functions are available through invocations
in the input text. Each such invocation contains a period as the
first character on the line, followed immediately by the name of
the function, followed by any arguments. A typical invocation might
be
	.play demo.go /^func show/,/^}/
(except that the ".play" must be at the beginning of the line and
not be indented like this.)

Here follows a description of the functions:

code:

Injects program source into the output by extracting code from files
and injecting them as HTML-escaped <pre> blocks.  The argument is
a file name followed by an optional address that specifies what
section of the file to display. The address syntax is similar in
its simplest form to that of ed, but comes from sam and is more
general. See
	http://plan9.bell-labs.com/sys/doc/sam/sam.html Table II
for full details. The displayed block is always rounded out to a
full line at both ends.

If no pattern is present, the entire file is displayed.

Any line in the program that ends with the four characters
	OMIT
is deleted from the source before inclusion, making it easy
to write things like
	.code test.go /START OMIT/,/END OMIT/
to find snippets like this
	tedious_code = boring_function()
	// START OMIT
	interesting_code = fascinating_function()
	// END OMIT
and see only this:
	interesting_code = fascinating_function()

Also, inside the displayed text a line that ends
	// HL
will be highlighted in the display; the 'h' key in the browser will
toggle extra emphasis of any highlighted lines. A highlighting mark
may have a suffix word, such as
	// HLxxx
Such highlights are enabled only if the code invocation ends with
"HL" followed by the word:
	.code test.go /^type Foo/,/^}/ HLxxx

play:

The function "play" is the same as "code" but puts a button
on the displayed source so the program can be run from the browser.
Although only the selected text is shown, all the source is included
in the HTML output so it can be presented to the compiler.

link:

Create a hyperlink. The syntax is 1 or 2 space-separated arguments.
The first argument is always the HTTP URL.  If there is a second
argument, it is the text label to display for this link.

	.link http://golang.org golang.org

image:

The template uses the function "image" to inject picture files.

The syntax is simple: 1 or 3 space-separated arguments.
The first argument is always the file name.
If there are more arguments, they are the height and width;
both must be present.

	.image images/betsy.jpg 100 200

html:

The function html includes the contents of the specified file as
unescaped HTML. This is useful for including custom HTML elements
that cannot be created using only the slide format.
It is your responsibilty to make sure the included HTML is valid and safe.

	.html file.html

*/
package main
