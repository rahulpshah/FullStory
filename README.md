Hi FullStory Team,

Thank you for giving me this wonderful opportunity to take such an interesting coding challenge.
I implemented 2 of the 3 questions which were mentioned in the challenge, as I found both of the
questions equally interesting.

In the repository, you will find a main file and mapper folder. The mapper folder has the implementation
of the hash based data structure. Associated tests are also included respectively.

Some of the points I considered for the hash map datastructure:
1. I wanted to created a map structure which allows custom hash function which is currently not
supported by golang
2. I have tried to optimize the data structure for space, preallocating the memory beforehand rather than
   using slice at startup.
3. Some of the other optimization can be made based on the access pattern of the users, for e.g. swapping
   the latest queried element in map to the start of the list in their bucket, so that next time it is available
   faster.

Regarding the elegance question, I believe the code is elegant if it is stateless, maintainable and readable.
Sometimes elegance comes at cost of performance, so one needs to make the tradeoff when to go for elegance and when
to go for performance.

Would love to hear back your thoughts wherever there is an room for improvement.

Looking forward to hearing from you.
