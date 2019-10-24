
"""
Function: emailFilter
Input: 
    friendList: a list of friends' email addresses
    allEmails: a list of all emails in the inbox
Output:
    the number of emails from friends
"""
def emailFilter(friendList, allEmails):
    count = 0 # represents the number of emails from friends
    for email in allEmails:
        if email in friendList:
            count += 1
    return count

friendList = ['friend@catena-technologies.com', 'bff@uni.edu.sg', 
              'provost@uni.edu.sg', 'john@doe.com']
allEmails = ['enemy@uni.edu.sg', 'bff@uni.edu.sg', 'evil@spam.com', 
             'bff@uni.edu.sg', 'friend@catena-technologies.com']

# Below is f-string in Python 3. It s good to know!
print(f'There are {emailFilter(friendList, allEmails)} emails from friends in Inbox\n')

# If you do not use f-string, here is traditional way
print('There are {0} emails from friends in Inbox\n'.format(emailFilter(friendList, allEmails)))