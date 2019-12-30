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
    <h1>Hello<h2> World</h2></h1>
</body>
</html>
"""


# def match_html_tags(start_tags, end_tags):
#
#     while not end_tags.is_empty():
#
#         if start_tags.peek() == end_tags.peek():
#             end_tags.pop()
#             start_tags.pop()
#         else:
#             return False


def check_html_match(html_input_content):
    token_location = 0
    start_tags_stack = Stack()
    end_tags_stack = Stack()

    for token in html_input_content:
        token_location += 1

        if token == '<':

            if html_input_content[token_location] == '/':
                reader = 1

                while html_input_content[token_location + reader] != '>':
                    end_tags_stack.push(html_input_content[token_location + reader])
                    reader += 1

                while not end_tags_stack.is_empty():
                    if start_tags_stack.peek() == end_tags_stack.peek():
                        start_tags_stack.pop()
                        end_tags_stack.pop()

                    else:
                        return False

            elif html_input_content[token_location] == '!':
                continue

            else:
                reader = 0

                while html_input_content[token_location + reader] != ' ' \
                        and html_input_content[token_location + reader] != '>':
                    start_tags_stack.push(html_input_content[token_location + reader])
                    reader += 1

    return start_tags_stack.is_empty()


print(check_html_match(html_content))



