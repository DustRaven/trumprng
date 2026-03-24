package trumprng

// quotes enthält die eingebettete Datenbank der Trump-Zitate.
// Jedes Zitat dient als potenzieller Seed-Wert für den PRNG.
var quotes = []Quote{
	// --- Selbsteinschätzung ---
	{
		Text:    "I have the best words. Nobody has better words than me.",
		Context: "Campaign rally",
		Year:    2015,
	},
	{
		Text:    "Sorry losers and haters, but my IQ is one of the highest — and you all know it!",
		Context: "Twitter",
		Year:    2013,
	},
	{
		Text:    "I'm like a very smart person. I know what I'm doing and I listen to everybody.",
		Context: "Press conference",
		Year:    2017,
	},
	{
		Text:    "I have a gut, and my gut tells me more sometimes than anybody else's brain can ever tell me.",
		Context: "Washington Post interview",
		Year:    2018,
	},
	{
		Text:    "I think I am actually humble. I think I'm much more humble than you would understand.",
		Context: "60 Minutes interview",
		Year:    2015,
	},
	{
		Text:    "I'm the most successful person ever to run for the presidency, by far.",
		Context: "NBC News",
		Year:    2015,
	},
	{
		Text:    "Nobody knows more about technology than me, believe me. Nobody.",
		Context: "Press gaggle",
		Year:    2015,
	},
	{
		Text:    "Nobody knows more about debt than me. I'm like the king of debt.",
		Context: "CNN interview",
		Year:    2016,
	},
	{
		Text:    "Nobody knows more about taxes than I do, maybe in the history of the world.",
		Context: "Press conference",
		Year:    2016,
	},
	{
		Text:    "I know more about ISIS than the generals do, believe me.",
		Context: "NBC News",
		Year:    2015,
	},
	{
		Text:    "I know more about wind than you do. Wind is a disaster.",
		Context: "CPAC speech",
		Year:    2019,
	},
	{
		Text:    "I know more about drones than anybody. I know about every form of safety that you can have.",
		Context: "Fox News",
		Year:    2013,
	},
	{
		Text:    "Nobody knows more about construction than I do. I built a lot of buildings.",
		Context: "Campaign speech",
		Year:    2016,
	},
	{
		Text:    "I know more about the big bills than any president that's ever been in office.",
		Context: "White House remarks",
		Year:    2019,
	},

	// --- Philosophie & Weltanschauung ---
	{
		Text:    "I could stand in the middle of Fifth Avenue and shoot somebody and wouldn't lose any voters.",
		Context: "Iowa rally",
		Year:    2016,
	},
	{
		Text:    "The concept of global warming was created by and for the Chinese in order to make U.S. manufacturing non-competitive.",
		Context: "Twitter",
		Year:    2012,
	},
	{
		Text:    "I think if this country gets any kinder or gentler, it's literally going to cease to exist.",
		Context: "Playboy interview",
		Year:    1990,
	},
	{
		Text:    "My whole life is about winning. I don't lose often. I almost never lose.",
		Context: "Campaign trail",
		Year:    2016,
	},
	{
		Text:    "The point is that you can't be too greedy.",
		Context: "The Art of the Deal",
		Year:    1987,
	},
	{
		Text:    "When somebody challenges you, fight back. Be brutal, be tough.",
		Context: "Think Big",
		Year:    2007,
	},
	{
		Text:    "I'm not running for office. I don't have to be politically correct.",
		Context: "Various speeches",
		Year:    2011,
	},

	// --- Naturphänomene ---
	{
		Text:    "Wind is a disaster for birds. Many, many birds are killed every year. And it's considered a great thing.",
		Context: "CPAC speech",
		Year:    2019,
	},
	{
		Text:    "The windmills are driving everybody crazy. They're killing all the birds. They're killing everything.",
		Context: "Remarks at the White House",
		Year:    2019,
	},
	{
		Text:    "It's freezing in New York — where the hell is global warming when you need it?",
		Context: "Twitter",
		Year:    2014,
	},
	{
		Text:    "In the old days, they didn't have air conditioning. And people were very productive.",
		Context: "White House remarks",
		Year:    2019,
	},

	// --- Beziehungen & Diplomatie ---
	{
		Text:    "Number one, I have great chemistry with Kim Jong Un. Great chemistry.",
		Context: "Press conference",
		Year:    2019,
	},
	{
		Text:    "Putin is fine. He's fine. We get along great.",
		Context: "Campaign rally",
		Year:    2016,
	},
	{
		Text:    "I've done more for women than almost any other presidential candidate in the history of this country.",
		Context: "Press conference",
		Year:    2016,
	},
	{
		Text:    "I have a great relationship with the blacks. I've always had a great relationship with the blacks.",
		Context: "Albany radio interview",
		Year:    2011,
	},

	// --- Wirtschaft & Handel ---
	{
		Text:    "Trade wars are good, and easy to win.",
		Context: "Twitter",
		Year:    2018,
	},
	{
		Text:    "I love the old days. You know what they used to do to guys like that when they were in a place like this? They'd be carried out on a stretcher, folks.",
		Context: "Campaign rally",
		Year:    2016,
	},
	{
		Text:    "Part of the beauty of me is that I'm very rich.",
		Context: "Good Morning America",
		Year:    2011,
	},
	{
		Text:    "If I ran for president, I'd win. I have the best product.",
		Context: "Wired interview",
		Year:    2012,
	},

	// --- Gesundheit & Medizin ---
	{
		Text:    "And then I see the disinfectant, where it knocks it out in a minute. Is there a way we can do something like that by injection inside?",
		Context: "White House briefing",
		Year:    2020,
	},
	{
		Text:    "Coronavirus. Wow. I never heard of it before. Now suddenly it's all people are talking about.",
		Context: "Campaign rally",
		Year:    2020,
	},
	{
		Text:    "I know that human being and fish can coexist peacefully.",
		Context: "Saginaw, Michigan",
		Year:    2000,
	},

	// --- Medien & Kommunikation ---
	{
		Text:    "I'm the most fabulous whiner. I do whine because I want to win.",
		Context: "CNN State of the Union",
		Year:    2015,
	},
	{
		Text:    "The fake news is not my enemy. It is the enemy of the American people.",
		Context: "CPAC speech",
		Year:    2017,
	},
	{
		Text:    "My Twitter has become so powerful that I can actually make my enemies tell the truth.",
		Context: "Twitter",
		Year:    2012,
	},
	{
		Text:    "You know what uranium is, right? It's this thing called nuclear weapons and other things. Like lots of things are done with uranium including some bad things.",
		Context: "Press conference",
		Year:    2016,
	},

	// --- Architektur & Ästhetik ---
	{
		Text:    "I am going to build a great, great wall on our southern border, and I will make Mexico pay for that wall. Mark my words.",
		Context: "Presidential announcement speech",
		Year:    2015,
	},
	{
		Text:    "The beauty of me is that I'm very rich.",
		Context: "Good Morning America",
		Year:    2011,
	},
	{
		Text:    "My fingers are long and beautiful, as, it has been well documented, are various other parts of my body.",
		Context: "New York Post",
		Year:    2011,
	},

	// --- 2. Amtszeit: Außenpolitik & Territorialphantasien (2025–2026) ---
	{
		Text:    "Canada should become our Cherished 51st State.",
		Context: "Truth Social",
		Year:    2025,
	},
	{
		Text:    "I don't want to spend hundreds of billions of dollars on supporting a country unless that country is a state.",
		Context: "Air Force One press gaggle",
		Year:    2025,
	},
	{
		Text:    "I love Canada. I have so many friends up in Canada. And they like me. But Canada has been taking advantage of the United States for years.",
		Context: "Air Force One press gaggle",
		Year:    2025,
	},
	{
		Text:    "We need Greenland for strategic national security and international security. This enormous, unsecured island is actually part of North America. That's our territory.",
		Context: "Davos speech",
		Year:    2026,
	},
	{
		Text:    "They have a choice. You can say yes, and we will be very appreciative, or you can say no and we will remember.",
		Context: "Davos, on Greenland",
		Year:    2026,
	},
	{
		Text:    "Denmark's defense of Greenland consists of two dogsleds.",
		Context: "Truth Social",
		Year:    2026,
	},
	{
		Text:    "The European Union was formed in order to screw the United States.",
		Context: "First Cabinet meeting",
		Year:    2025,
	},
	{
		Text:    "Look, the Panama Canal is vital to our country. It's being operated by China. And we gave the Panama Canal to Panama.",
		Context: "Mar-a-Lago press conference",
		Year:    2025,
	},
	{
		Text:    "We're going to make Gaza the Riviera of the Middle East.",
		Context: "Press conference with Netanyahu",
		Year:    2025,
	},
	{
		Text:    "I am hearing that the people of Greenland are MAGA.",
		Context: "Truth Social",
		Year:    2025,
	},

	// --- 2. Amtszeit: Wirtschaft & Tarife ---
	{
		Text:    "Trade wars are good, and easy to win. We've proven that now.",
		Context: "White House remarks",
		Year:    2025,
	},
	{
		Text:    "The word affordability is a Democrat scam.",
		Context: "Cabinet meeting",
		Year:    2025,
	},
	{
		Text:    "They came up with a new word — affordability.",
		Context: "US-Saudi Investment Forum, Kennedy Center",
		Year:    2025,
	},
	{
		Text:    "Our stock market dip is peanuts. That stock market is going to be doubled. The Dow is going to hit 100,000 in a relatively short period of time.",
		Context: "Davos speech",
		Year:    2026,
	},
	{
		Text:    "With tariffs, we've radically reduced our ballooning trade deficit, which was the largest in world history. We were losing more than a trillion dollars every single year and it was just wasted.",
		Context: "Davos speech",
		Year:    2026,
	},

	// --- 2. Amtszeit: Selbstbild & Sendungsbewusstsein ---
	{
		Text:    "I was saved by God to make America great again.",
		Context: "First day back in office, address to Congress",
		Year:    2025,
	},
	{
		Text:    "These countries are calling us up, kissing my ass.",
		Context: "Republican conference",
		Year:    2025,
	},
	{
		Text:    "My biggest surprise is I thought it would take more than a year, maybe like a year and one month. But it's happened very quickly.",
		Context: "Davos speech, on economic recovery",
		Year:    2026,
	},
	{
		Text:    "It's great to be in this beautiful Oval Office. It never looked so nice. We're getting great compliments.",
		Context: "FBI press conference",
		Year:    2025,
	},
	{
		Text:    "I really predict that. We're gonna start working together in healthcare.",
		Context: "Congressional Ball speech",
		Year:    2025,
	},

	// --- 2. Amtszeit: Geopolitik & Sonstiges ---
	{
		Text:    "Something that you could say 3,000 years, if you look at it in certain ways, or you could say centuries. But this is a deal that incredibly everyone just came together.",
		Context: "Press conference on Middle East peace",
		Year:    2025,
	},
	{
		Text:    "Even the enemies came out and totally endorsed it, everybody. I've never seen anything like it. Nobody's seen anything like it.",
		Context: "White House remarks on Middle East",
		Year:    2025,
	},
	{
		Text:    "Certain places in Europe are not recognisable, frankly, any more. I want to see Europe go good, but it's not heading in the right direction.",
		Context: "Davos speech",
		Year:    2026,
	},
	{
		Text:    "It's horrible what they've done to themselves.",
		Context: "Davos speech, über Europa",
		Year:    2026,
	},
	{
		Text:    "I no longer feel an obligation to think purely of Peace.",
		Context: "Message to NATO leaders, on Greenland",
		Year:    2026,
	},

	// --- Klassiker ---
	{
		Text:    "We will have so much winning if I get elected that you may get bored with winning.",
		Context: "Campaign rally",
		Year:    2015,
	},
	{
		Text:    "Nobody builds walls better than me, believe me, and I'll build them very inexpensively.",
		Context: "Presidential announcement",
		Year:    2015,
	},
	{
		Text:    "I will be the greatest jobs president that God ever created. I tell you that.",
		Context: "Presidential announcement",
		Year:    2015,
	},
	{
		Text:    "I have never seen a thin person drinking Diet Coke.",
		Context: "Twitter",
		Year:    2012,
	},
	{
		Text:    "You know, it really doesn't matter what the media writes as long as you've got a young and beautiful piece of ass.",
		Context: "Esquire",
		Year:    1991,
	},
}
