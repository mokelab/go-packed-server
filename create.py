# -*- coding: utf-8 -*-

import os
import base64

def get_content_type(file):
    if file.endswith(".html"):
        return "text/html"
    return "text/plain"

def write_content(folder, root):
    files = os.listdir(folder)
    for file in files:
        full_path = "%s/%s" % (folder, file)
        relative_path = "%s%s" % (root, file)
        if os.path.isdir(full_path):
            write_content(full_path, "%s/" % (relative_path))
            continue
        content_type = get_content_type(file)
        f = open(full_path, 'rt').read()
        enc_file = base64.b64encode(f)
        print """\tc.Entry[\"%s\"] = Entry{
    ContentType: "%s",
    src: "%s",
    }
""" % (relative_path, content_type, enc_file)

files = os.listdir("./contents")

print """package data

func Create() *Cache {
	c := &Cache{
		Entry: make(map[string]Entry),
	}
"""

write_content("./contents", "/")

print "return c}"
