# Back-end Coding Challenge

Welcome to this back-end Coding Challenge! We will ask you to complete the following exercise as part of our Interview process. It aims at assessing your coding skills proficiency, and your problem-solving mindset. It should take you no more than 2h.

You can submit this article as a public GitHub repository with the exercise and a short README explaining how to run it. 

You can write your code in **Go** (preferred), **JavaScript** or **Python**.

We will assess:
- If your solution works as expected
- The way you structure the code and its cleanliness
- Adherence to established standards when building API services
- The way you introduce code into your Git history

# Back-end Coding Challenge

Welcome to this back-end Coding Challenge! We will ask you to complete the following exercise as part of our Interview process. It aims at assessing your coding skills proficiency, and your problem-solving mindset. It should take you no more than 2h.

You can submit this article as a public GitHub repository with the exercise and a short README explaining how to run it. 

You can write your code in **Go** (preferred), **JavaScript** or **Python**.

````
User: {
	id: int
	name: string
	createdAt: date
}

Action: {
	id: int
	type: string
	userId: int       // The ID of the User who performed this action
	targetUser: int   // Supplied when "REFER_USER" action type
	createdAt: date
}
````

### Goal

Deliver a **very simple web server** capable of querying the relevant API endpoints described below.

### Additional Info

There is **no need to implement the database layer**, reading the file in memory at startup is sufficient. Use the **two JSON files** below containing a database sample for this exercise.

### 1) Write an API endpoint to fetch a User given its ID

Example response:

````
{
	id: "1234"
	name: "John Doe"
	createdAt: "2022-04-14T11:12:22.758Z"
}
````

### 2) Write an API endpoint to get the total number of actions of a User given its ID

Example response:

````
{
	count: 100
}
````

### 3) Write an API endpoint to get the break-down of all possible next actions given an action Type

Users perform different actions throughout the day when using our product. Right after doing an action A, users could perform action B, C, or A again for instance. Write an endpoint that gives a probability break-down of what actions our users typically do after doing action A, based on the entire Actions database.

Example response:

````
{
	ADD_TO_CRM: 0.70
	REFER_USER: 0.20
	VIEW_CONVERSATION: 0.10
}
````

### 4) Write an efficient API endpoint to get the “Referral Index” of all the Users. Discuss complexity.

Users have the ability to refer other users (i.e. inviting them to use the product as well).

When doing so, an activity with type `REFER_USER` is created by the existing user, and the ID of the new invited user is stored in the  `targetUser` attribute.

We can therefore compute the **Referral Index** for a given user as the total number of individual users invited directly or indirectly by this user (users invited by the given user can also invite new users, etc...)

Assume a user can be invited only ONCE.

Example response:

````
{
	1: 3 // UserID: Referral index
	2: 0
	3: 7
	...
}
````