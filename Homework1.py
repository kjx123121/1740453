"""
Jxkang
12/30/19

Project requirement:
Receive a string of HTML codes and match starting tags with ending tags return as boolean
if starting tag match with ending tags return True otherwise False

"""

from listRearRepresentStackTop import Stack

html_content = """
<!DOCTYPE html>
<html lang="en">
<head>
    <title>Title</title>
</head>
<body>

    <h1>Hello World</h1>
</body>
</html>
"""

html_content2 = """
<!DOCTYPE html>
<html lang="en">
<head>
    <title>Title</title>
</head>
<body>
    <h1>Hello<h2> World</h1></h2>
</body>
</html>
"""


def check_html_match(string_of_html):
    next_token_location = 0
    start_tags_stack = Stack()
    end_tags_stack = Stack()

    # Read token from string of HTML find all '<' since all tags start with '<'
    for token in string_of_html:
        next_token_location += 1

        if token == '<':

            # If there is '/' which means this is a ending tags
            if string_of_html[next_token_location] == '/':
                tags_char_location = 1

                # Push all ending tags into end_tags_stack
                while string_of_html[next_token_location + tags_char_location] != '>':     
                    end_tags_stack.push(string_of_html[next_token_location + tags_char_location])
                    tags_char_location += 1

                # For each time we finished with inserting single end tags compare with start tags
                while not end_tags_stack.is_empty():
                    
                    if start_tags_stack.peek() == end_tags_stack.peek():
                        start_tags_stack.pop()
                        end_tags_stack.pop()

                    else:
                        return False

            elif string_of_html[next_token_location] == '!':
                continue

            # Push all starting tags into start_tags_stack
            else:
                tags_char_location = 0

                while string_of_html[next_token_location + tags_char_location] != ' ' \
                        and string_of_html[next_token_location + tags_char_location] != '>':
                    start_tags_stack.push(string_of_html[next_token_location + tags_char_location])
                    tags_char_location += 1

    return start_tags_stack.is_empty()


print(check_html_match(html_content2))



