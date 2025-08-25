I am writing this post while I am still building this blog. I was getting sick of figuring out HTML templates, so I thought I would build a markdown to HTML parser, which I am going to need for this blog anyway. This is the first entrant for my coding across the alphabet challenge, although, this one is out of order. Anyways, P for Parser it is!

I haven't actually written the parser yet though. I thought it would be fun to write about it as I built it, so I can use the features in this post as I implement them. Currently, this is just normal text in a <p> tag, nothing fancy.

I did a bit of research into popular parsing methods, but it turns out that they are mostly overkill for simple markdown. Parsing line by line and handling each rule seperately should be fine. Very helpfully for me, each rule is defined in commonmark.org, which hosts the spec for different versions and other helpful info. It also includes test cases! So now I don't have to come up with tests. They come in a JSON format, so first I will write some logic to parse the tests and make them compatible with Go's testing system.

The format of a test is as follows:

{
"markdown": "# foo\n## foo\n### foo\n#### foo\n##### foo\n###### foo\n",
"html": "<h1>foo</h1>\n<h2>foo</h2>\n<h3>foo</h3>\n<h4>foo</h4>\n<h5>foo</h5>\n<h6>foo</h6>\n",
"example": 62,
"start_line": 1112,
"end_line": 1126,
"section": "ATX headings"
},
