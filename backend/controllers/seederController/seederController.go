package seedercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/Itsjoses/sebook-be/database"
	"github.com/Itsjoses/sebook-be/models"
	"github.com/gin-gonic/gin"
)

func BookSeeder(c *gin.Context) {
	jsonData := `[
		{
		  "Title": "Alice's Adventures in Wonderland",
		  "Cover": 10527843,
		  "GenreID": 1,
		  "Price": 104
		},
		{
		  "Title": "Adventures of Huckleberry Finn",
		  "Cover": 8157718,
		  "GenreID": 1,
		  "Price": 95
		},
		{
		  "Title": "The Adventures of Tom Sawyer",
		  "Cover": 12043351,
		  "GenreID": 1,
		  "Price": 109
		},
		{
		  "Title": "The Adventures of Sherlock Holmes [12 stories]",
		  "Cover": 6717853,
		  "GenreID": 4,
		  "Price": 78
		},
		{
		  "Title": "Alice's Adventures in Wonderland / Through the Looking Glass",
		  "Cover": 8595966,
		  "GenreID": 2,
		  "Price": 71
		},
		{
		  "Title": "The Merry Adventures of Robin Hood of Great Renown in Nottinghamshire",
		  "Cover": 5913135,
		  "GenreID": 3,
		  "Price": 123
		},
		{
		  "Title": "The Adventures of Captain Underpants",
		  "Cover": 12625278,
		  "GenreID": 4,
		  "Price": 107
		},
		{
		  "Title": "The Life and Adventures of Nicholas Nickleby",
		  "Cover": 9283905,
		  "GenreID": 2,
		  "Price": 108
		},
		{
		  "Title": "The Celestine Prophecy - An Adventure",
		  "Cover": 8749217,
		  "GenreID": 3,
		  "Price": 85
		},
		{
		  "Title": "The Complete Life and Adventures of Santa Claus",
		  "Cover": 1979059,
		  "GenreID": 4,
		  "Price": 146
		},
		{
		  "Title": "The Story of the Treasure Seekers being the adventures of the Bastable children in search of a fortune",
		  "Cover": 13241364,
		  "GenreID": 3,
		  "Price": 70
		},
		{
		  "Title": "The Amazing Adventures of Kavalier & Clay",
		  "Cover": 12745205,
		  "GenreID": 1,
		  "Price": 184
		},
		{
		  "Title": "The History of the Adventures of Joseph Andrews",
		  "Cover": 5955750,
		  "GenreID": 4,
		  "Price": 67
		},
		{
		  "Title": "Novels (Adventures of Huckleberry Finn / Adventures of Tom Sawyer)",
		  "Cover": 12374727,
		  "GenreID": 3,
		  "Price": 166
		},
		{
		  "Title": "The Adventures of Gerard",
		  "Cover": 8243324,
		  "GenreID": 4,
		  "Price": 72
		},
		{
		  "Title": "The Complete Adventures of Charlie and Willy Wonka (Charlie and the Chocolate Factory / Charlie and the Great Glass Elevator)",
		  "Cover": 10606361,
		  "GenreID": 1,
		  "Price": 80
		},
		{
		  "Title": "Salute to Adventurers",
		  "Cover": 10710393,
		  "GenreID": 4,
		  "Price": 176
		},
		{
		  "Title": "The little white bird, or, Adventures in Kensington gardens",
		  "Cover": 2808666,
		  "GenreID": 3,
		  "Price": 102
		},
		{
		  "Title": "The adventures of Super Diaper Baby",
		  "Cover": 277462,
		  "GenreID": 2,
		  "Price": 145
		},
		{
		  "Title": "Arthur Writes a Story (Arthur Adventure Series)",
		  "Cover": 188077,
		  "GenreID": 2,
		  "Price": 139
		},
		{
		  "Title": "The Lottery, or The Adventures of James Harris",
		  "Cover": 9272736,
		  "GenreID": 2,
		  "Price": 122
		},
		{
		  "Title": "Five Go Adventuring Again",
		  "Cover": 11627305,
		  "GenreID": 4,
		  "Price": 129
		},
		{
		  "Title": "Business adventures",
		  "Cover": 8231778,
		  "GenreID": 2,
		  "Price": 110
		},
		{
		  "Title": "Strange Adventures of Mr. Middleton Illustrated",
		  "Cover": 11534189,
		  "GenreID": 4,
		  "Price": 189
		},
		{
		  "Title": "The Adventure of the Bruce-Partington Plans",
		  "Cover": 2037348,
		  "GenreID": 2,
		  "Price": 82
		},
		{
		  "Title": "What the dog saw and other adventure stories",
		  "Cover": 8260666,
		  "GenreID": 3,
		  "Price": 188
		},
		{
		  "Title": "The Adventure of the Christmas Pudding",
		  "Cover": 12855353,
		  "GenreID": 1,
		  "Price": 158
		},
		{
		  "Title": "Star Wars - From the Adventures of Luke Skywalker",
		  "Cover": 11073173,
		  "GenreID": 2,
		  "Price": 84
		},
		{
		  "Title": "The Adventure of the Speckled Band",
		  "Cover": 8517526,
		  "GenreID": 3,
		  "Price": 126
		},
		{
		  "Title": "Choose Your Own Adventure - Journey Under The Sea",
		  "Cover": 3062257,
		  "GenreID": 1,
		  "Price": 153
		},
		{
		  "Title": "Adventures of Roderick Random",
		  "Cover": 8245241,
		  "GenreID": 2,
		  "Price": 105
		},
		{
		  "Title": "The Adventures of Caleb Williams",
		  "Cover": 6002901,
		  "GenreID": 3,
		  "Price": 196
		},
		{
		  "Title": "The adventures of Peregrine Pickle",
		  "Cover": 5735518,
		  "GenreID": 3,
		  "Price": 144
		},
		{
		  "Title": "The Adventures of Odysseus And Tales of Troy",
		  "Cover": 1754536,
		  "GenreID": 4,
		  "Price": 95
		},
		{
		  "Title": "Thea Stilton and the Cherry Blossom Adventure (Thea Stilton 6)",
		  "Cover": 14420756,
		  "GenreID": 4,
		  "Price": 57
		},
		{
		  "Title": "The adventures of Augie March",
		  "Cover": 400376,
		  "GenreID": 1,
		  "Price": 175
		},
		{
		  "Title": "Arthur's New Puppy (Arthur Adventure Series)",
		  "Cover": 188015,
		  "GenreID": 2,
		  "Price": 57
		},
		{
		  "Title": "The Island of Adventure",
		  "Cover": 6256759,
		  "GenreID": 3,
		  "Price": 161
		},
		{
		  "Title": "The Adventures of Beekle",
		  "Cover": 7337423,
		  "GenreID": 4,
		  "Price": 176
		},
		{
		  "Title": "Draw Your Own Adventure, Blank Comic Book",
		  "Cover": 13182376,
		  "GenreID": 4,
		  "Price": 193
		},
		{ "Title": "Adventure", "Cover": 8233626, "GenreID": 2, "Price": 198 },
		{
		  "Title": "The Viking Adventure",
		  "Cover": 6901403,
		  "GenreID": 2,
		  "Price": 70
		},
		{
		  "Title": "Escape on Venus (Carson Napier Adventures #4) (Ace SF Classic, 21561)",
		  "Cover": 14463614,
		  "GenreID": 2,
		  "Price": 123
		},
		{
		  "Title": "Alex's Adventures In Numberland",
		  "Cover": 7544307,
		  "GenreID": 1,
		  "Price": 129
		},
		{
		  "Title": "The Adventures of Tintin",
		  "Cover": 6979825,
		  "GenreID": 3,
		  "Price": 107
		},
		{
		  "Title": "Adventures in the screen trade",
		  "Cover": 4277474,
		  "GenreID": 4,
		  "Price": 96
		},
		{
		  "Title": "Wonderful adventures of Mrs. Seacole in many lands",
		  "Cover": 1124260,
		  "GenreID": 2,
		  "Price": 168
		},
		{
		  "Title": "The Adventures of Robin Hood",
		  "Cover": 418176,
		  "GenreID": 3,
		  "Price": 137
		},
		{
		  "Title": "The Adventure of Wisteria Lodge",
		  "Cover": 9988345,
		  "GenreID": 1,
		  "Price": 105
		},
		{
		  "Title": "Pokemon Adventures",
		  "Cover": 4838442,
		  "GenreID": 1,
		  "Price": 184
		},
		{
		  "Title": "Roblox Top Adventure Games",
		  "Cover": 8814513,
		  "GenreID": 3,
		  "Price": 162
		},
		{
		  "Title": "Raffles, Further Adventures of the Amateur Cracksman",
		  "Cover": 5962788,
		  "GenreID": 2,
		  "Price": 185
		},
		{
		  "Title": "Secret Seven Adventure",
		  "Cover": 5076683,
		  "GenreID": 2,
		  "Price": 160
		},
		{
		  "Title": "Choose Your Own Adventure - The Cave of Time",
		  "Cover": 4672785,
		  "GenreID": 4,
		  "Price": 168
		},
		{
		  "Title": "Our Mr. Wrenn (The Romantic Adventures of A Gentle Man)",
		  "Cover": 1760451,
		  "GenreID": 4,
		  "Price": 118
		},
		{
		  "Title": "The original adventures of Hank the Cowdog",
		  "Cover": 4062524,
		  "GenreID": 4,
		  "Price": 115
		},
		{
		  "Title": "The Mountain of Adventure",
		  "Cover": 8237543,
		  "GenreID": 3,
		  "Price": 167
		},
		{
		  "Title": "The Adventures of Tom Bombadil and Other Verses from the Red Book",
		  "Cover": 8220109,
		  "GenreID": 2,
		  "Price": 143
		},
		{
		  "Title": "The Works of Edgar Allan Poe in Five Volumes - Volume One (Balloon-Hoax / Four Beasts in One / Gold-Bug / Ms. Found in a Bottle / Murders in the Rue Morgue / Mystery of Marie Roget / Oval Portrait / Unparalleled Adventure of One Hans Pfaall)",
		  "Cover": 11657982,
		  "GenreID": 1,
		  "Price": 107
		},
		{
		  "Title": "The Akhenaten Adventure",
		  "Cover": 1315938,
		  "GenreID": 4,
		  "Price": 160
		},
		{
		  "Title": "Alice's Adventures in Wonderland / Through the Looking-Glass / The Hunting of the Snark",
		  "Cover": 10527372,
		  "GenreID": 1,
		  "Price": 146
		},
		{
		  "Title": "The Secret Seven and the Tree House Adventure",
		  "Cover": 9814280,
		  "GenreID": 1,
		  "Price": 132
		},
		{
		  "Title": "The Castle of Adventure",
		  "Cover": 1180283,
		  "GenreID": 1,
		  "Price": 135
		},
		{
		  "Title": "The Adventure of the Devil's Foot",
		  "Cover": 9005885,
		  "GenreID": 3,
		  "Price": 165
		},
		{
		  "Title": "The Sea of Adventure",
		  "Cover": 1180285,
		  "GenreID": 3,
		  "Price": 53
		},
		{
		  "Title": "Rowley Jefferson's Awesome Friendly Adventure",
		  "Cover": 10328978,
		  "GenreID": 4,
		  "Price": 80
		},
		{
		  "Title": "The splendid spur, being memoirs of the adventures of Mr. John Marvel, a servant of His late Majesty King Charles I., in the years 1642-3",
		  "Cover": 5584204,
		  "GenreID": 3,
		  "Price": 188
		},
		{
		  "Title": "Adventures in English literature",
		  "Cover": 6822832,
		  "GenreID": 4,
		  "Price": 104
		},
		{
		  "Title": "The Adventures of Harry Richmond",
		  "Cover": 1814182,
		  "GenreID": 4,
		  "Price": 52
		},
		{
		  "Title": "The Adventures of the Wishing-Chair",
		  "Cover": 1314652,
		  "GenreID": 1,
		  "Price": 158
		},
		{
		  "Title": "Five Fall into Adventure",
		  "Cover": 10640419,
		  "GenreID": 1,
		  "Price": 97
		},
		{
		  "Title": "Passionate Adventure",
		  "Cover": 9776208,
		  "GenreID": 4,
		  "Price": 60
		},
		{
		  "Title": "The Adventure of the Cardboard Box",
		  "Cover": 1801797,
		  "GenreID": 4,
		  "Price": 121
		},
		{
		  "Title": "The Black Mask: Further Adventure of the Amateur Cracksman",
		  "Cover": 10557056,
		  "GenreID": 2,
		  "Price": 110
		},
		{
		  "Title": "Eldorado Or Adventures In Path Of Empire",
		  "Cover": 2756244,
		  "GenreID": 4,
		  "Price": 134
		},
		{
		  "Title": "Trans-Himalaya; discoveries and adventures in Tibet",
		  "Cover": 3717305,
		  "GenreID": 2,
		  "Price": 180
		},
		{
		  "Title": "The adventures of Souza",
		  "Cover": 8494992,
		  "GenreID": 2,
		  "Price": 150
		},
		{
		  "Title": "Adventures in criticism",
		  "Cover": 5739468,
		  "GenreID": 1,
		  "Price": 52
		},
		{
		  "Title": "The Adventures of Reddy Fox",
		  "Cover": 310994,
		  "GenreID": 3,
		  "Price": 112
		},
		{
		  "Title": "The Adventures of Stainless Steel Rat",
		  "Cover": 4627656,
		  "GenreID": 1,
		  "Price": 128
		},
		{
		  "Title": "Pathfinders of the West: Being the Thrilling Story of the Adventures of the Men Who Discovered the Great Northwest",
		  "Cover": 2840468,
		  "GenreID": 1,
		  "Price": 116
		},
		{
		  "Title": "The Poseidon adventure",
		  "Cover": 111374,
		  "GenreID": 2,
		  "Price": 112
		},
		{
		  "Title": "Adventures of Charlie",
		  "Cover": 12436879,
		  "GenreID": 2,
		  "Price": 196
		},
		{
		  "Title": "Short Stories (Adventures of Sherlock Holmes / Memoirs of Sherlock Holmes [12 stories])",
		  "Cover": 9246437,
		  "GenreID": 3,
		  "Price": 155
		},
		{
		  "Title": "Choose Your Own Adventure - The Mystery of Chimney Rock",
		  "Cover": 6542695,
		  "GenreID": 1,
		  "Price": 135
		},
		{
		  "Title": "Works (Adventures of Sherlock Holmes / Hound of The Baskervilles /  Memoirs of Sherlock Holmes [12 stories] / Return of Sherlock Holmes)",
		  "Cover": 8349161,
		  "GenreID": 4,
		  "Price": 133
		},
		{
		  "Title": "Adventures of ideas",
		  "Cover": 6464980,
		  "GenreID": 2,
		  "Price": 119
		},
		{
		  "Title": "Goosebumps - My Hairiest Adventure",
		  "Cover": 384579,
		  "GenreID": 1,
		  "Price": 76
		},
		{
		  "Title": "Choose Your Own Adventure - Escape",
		  "Cover": 8311608,
		  "GenreID": 3,
		  "Price": 179
		},
		{
		  "Title": "Choose Your Own Adventure - Space And Beyond",
		  "Cover": 3062258,
		  "GenreID": 3,
		  "Price": 188
		},
		{
		  "Title": "The Adventures of Hajji Baba, of Ispahan",
		  "Cover": 10891496,
		  "GenreID": 1,
		  "Price": 153
		},
		{
		  "Title": "Adventures in Tandem Nursing",
		  "Cover": 703676,
		  "GenreID": 4,
		  "Price": 145
		},
		{
		  "Title": "The Adventures of Sally",
		  "Cover": 94126,
		  "GenreID": 4,
		  "Price": 143
		},
		{ "Title": "Amazon Adventure", "Cover": 70512, "GenreID": 1, "Price": 65 },
		{
		  "Title": "Two little savages ; being the adventures of two boys who lived as Indians and what they learned",
		  "Cover": 10934381,
		  "GenreID": 3,
		  "Price": 92
		},
		{
		  "Title": "Choose Your Own Adventure - Who Killed Harlowe Thrombey?",
		  "Cover": 6633222,
		  "GenreID": 2,
		  "Price": 62
		},
		{
		  "Title": "Choose Your Own Adventure - Mystery of the Maya",
		  "Cover": 8870229,
		  "GenreID": 3,
		  "Price": 60
		},
		{
		  "Title": "The Valley of Adventure",
		  "Cover": 1004201,
		  "GenreID": 4,
		  "Price": 68
		},
		{
		  "Title": "Lamb's adventures of Ulysses",
		  "Cover": 6121773,
		  "GenreID": 1,
		  "Price": 83
		},
		{
		  "Title": "The adventures of Johnny Chuck",
		  "Cover": 311667,
		  "GenreID": 3,
		  "Price": 110
		},
		{
		  "Title": "The Adventure of the Dying Detective",
		  "Cover": 1801798,
		  "GenreID": 1,
		  "Price": 99
		},
		{
		  "Title": "Arcadian Adventures with the Idle Rich",
		  "Cover": 11096873,
		  "GenreID": 4,
		  "Price": 130
		},
		{
		  "Title": "The Adventures of Hajji Baba of Ispahan",
		  "Cover": 875610,
		  "GenreID": 2,
		  "Price": 174
		},
		{
		  "Title": "The adventures of Odysseus",
		  "Cover": 885984,
		  "GenreID": 2,
		  "Price": 52
		},
		{
		  "Title": "Tennis Shoes Among the Nephites [Tennis shoes adventure series 1]",
		  "Cover": 4846353,
		  "GenreID": 3,
		  "Price": 113
		},
		{
		  "Title": "The surprising adventures of the magical monarch of Mo and his people",
		  "Cover": 1825119,
		  "GenreID": 2,
		  "Price": 192
		},
		{
		  "Title": "The Adventures of Captain Horn",
		  "Cover": 5735322,
		  "GenreID": 1,
		  "Price": 175
		},
		{
		  "Title": "The adventures of Captain Bonneville",
		  "Cover": 5947209,
		  "GenreID": 3,
		  "Price": 177
		},
		{
		  "Title": "Adventures Among Books",
		  "Cover": 4753088,
		  "GenreID": 3,
		  "Price": 58
		},
		{
		  "Title": "The life, adventures & piracies of the famous Captain Singleton",
		  "Cover": 11016225,
		  "GenreID": 4,
		  "Price": 138
		},
		{
		  "Title": "The River of Adventure",
		  "Cover": 11706748,
		  "GenreID": 2,
		  "Price": 156
		},
		{
		  "Title": "Choose Your Own Adventure - Vampire Express",
		  "Cover": 9828785,
		  "GenreID": 2,
		  "Price": 170
		},
		{
		  "Title": "The complete adventures of Curious George",
		  "Cover": 393639,
		  "GenreID": 3,
		  "Price": 121
		},
		{
		  "Title": "The Illryian Adventure",
		  "Cover": 5006006,
		  "GenreID": 3,
		  "Price": 196
		},
		{
		  "Title": "Houghton Mifflin Harcourt Journeys Reading Adventure",
		  "Cover": 14476164,
		  "GenreID": 1,
		  "Price": 169
		},
		{
		  "Title": "The Adventure Zone",
		  "Cover": 8214817,
		  "GenreID": 3,
		  "Price": 149
		},
		{
		  "Title": "The Mammoth Book of New Sherlock Holmes Adventures",
		  "Cover": 14570885,
		  "GenreID": 2,
		  "Price": 160
		},
		{
		  "Title": "The Adventures of Ook and Gluk, Kung Fu Cavemen from the Future",
		  "Cover": 6458473,
		  "GenreID": 4,
		  "Price": 75
		},
		{
		  "Title": "The adventures of Hajji Baba of Ispahan",
		  "Cover": 6121734,
		  "GenreID": 2,
		  "Price": 116
		},
		{
		  "Title": "Choose Your Own Adventure - Your Code Name is Jonah",
		  "Cover": 8871689,
		  "GenreID": 2,
		  "Price": 57
		},
		{
		  "Title": "The Ship of Adventure",
		  "Cover": 11713057,
		  "GenreID": 2,
		  "Price": 83
		},
		{
		  "Title": "Intellectual adventure of ancient man",
		  "Cover": 9469871,
		  "GenreID": 4,
		  "Price": 107
		},
		{
		  "Title": "Arcadian adventures with the idle rich",
		  "Cover": 5830537,
		  "GenreID": 1,
		  "Price": 171
		},
		{
		  "Title": "The adventures of Ferdinand Count Fathom",
		  "Cover": 7073816,
		  "GenreID": 4,
		  "Price": 82
		},
		{
		  "Title": "The adventures of Sir Launcelot Greaves",
		  "Cover": 2810249,
		  "GenreID": 2,
		  "Price": 200
		},
		{
		  "Title": "Star Trek Adventures - The New Voyages 2",
		  "Cover": 368865,
		  "GenreID": 1,
		  "Price": 108
		},
		{
		  "Title": "The 'Adventurers of England' on Hudson bay",
		  "Cover": 5929700,
		  "GenreID": 3,
		  "Price": 81
		},
		{
		  "Title": "Choose Your Own Adventure - The Abominable Snowman",
		  "Cover": 9720724,
		  "GenreID": 1,
		  "Price": 52
		},
		{
		  "Title": "Memories and Adventures",
		  "Cover": 9455077,
		  "GenreID": 2,
		  "Price": 84
		},
		{
		  "Title": "The adventures of Buster Bear",
		  "Cover": 311304,
		  "GenreID": 3,
		  "Price": 127
		},
		{
		  "Title": "The Adventures of the Princess and Mr. Whiffle - The Thing Beneath the Bed",
		  "Cover": 6691340,
		  "GenreID": 3,
		  "Price": 177
		},
		{
		  "Title": "Choose Your Own Adventure - Deadwood City",
		  "Cover": 8087376,
		  "GenreID": 1,
		  "Price": 199
		},
		{
		  "Title": "The Strange Adventures of Mr. Middleton",
		  "Cover": 9010733,
		  "GenreID": 4,
		  "Price": 186
		},
		{
		  "Title": "The Little Adventure",
		  "Cover": 12882443,
		  "GenreID": 2,
		  "Price": 96
		},
		{
		  "Title": "Adventures of Pinocchio",
		  "Cover": 12622151,
		  "GenreID": 1,
		  "Price": 74
		},
		{
		  "Title": "Choose Your Own Adventure - Return to the Cave of Time",
		  "Cover": 6631564,
		  "GenreID": 3,
		  "Price": 89
		},
		{
		  "Title": "The young Franc-tireurs and their adventures in the Franco-Prussian war",
		  "Cover": 5920365,
		  "GenreID": 2,
		  "Price": 87
		},
		{
		  "Title": "The adventures of Peter Cottontail",
		  "Cover": 310993,
		  "GenreID": 1,
		  "Price": 166
		},
		{
		  "Title": "Arthur's Halloween (Arthur Adventure Series)",
		  "Cover": 188056,
		  "GenreID": 4,
		  "Price": 161
		},
		{
		  "Title": "More Adventures of the Great Brain",
		  "Cover": 575419,
		  "GenreID": 3,
		  "Price": 161
		},
		{
		  "Title": "Choose Your Own Adventure - Beyond Escape!",
		  "Cover": 8870213,
		  "GenreID": 3,
		  "Price": 71
		},
		{
		  "Title": "An Awfully Big Adventure",
		  "Cover": 213270,
		  "GenreID": 3,
		  "Price": 68
		},
		{
		  "Title": "Pokemon Adventures, Volume 3",
		  "Cover": 813582,
		  "GenreID": 1,
		  "Price": 94
		},
		{
		  "Title": "The Improbable Adventures of Sherlock Holmes",
		  "Cover": 8358319,
		  "GenreID": 2,
		  "Price": 196
		},
		{
		  "Title": "Stories of American life and adventure",
		  "Cover": 5924702,
		  "GenreID": 3,
		  "Price": 73
		},
		{
		  "Title": "The Great Adventure (A Play of Fancy)",
		  "Cover": 1862919,
		  "GenreID": 1,
		  "Price": 87
		},
		{
		  "Title": "Adventures and Letters of Richard Harding Davis",
		  "Cover": 2761193,
		  "GenreID": 3,
		  "Price": 87
		},
		{
		  "Title": "Love's Wild Desire:(Love and Adventure Collection #2)",
		  "Cover": 506448,
		  "GenreID": 1,
		  "Price": 112
		},
		{
		  "Title": "The Adventures of Ulysses",
		  "Cover": 383767,
		  "GenreID": 1,
		  "Price": 100
		},
		{
		  "Title": "The adventures of Bert",
		  "Cover": 223660,
		  "GenreID": 4,
		  "Price": 162
		},
		{
		  "Title": "Incredible Adventures (Lovecraft's Library)",
		  "Cover": 740744,
		  "GenreID": 2,
		  "Price": 140
		},
		{
		  "Title": "Baby-Sitters' island adventure. (Baby-Sitters Club SUPER SPECIAL no.4)",
		  "Cover": 383759,
		  "GenreID": 4,
		  "Price": 185
		},
		{
		  "Title": "Vampires Don't Wear Polka Dots (Adventures of the Bailey School Kids, 1)",
		  "Cover": 6694607,
		  "GenreID": 3,
		  "Price": 136
		},
		{
		  "Title": "The Adventure of the Blue Carbuncle",
		  "Cover": 767108,
		  "GenreID": 3,
		  "Price": 53
		},
		{
		  "Title": "The Adventures of Elizabeth in Rügen",
		  "Cover": 6121732,
		  "GenreID": 4,
		  "Price": 121
		},
		{
		  "Title": "Star Trek - Enterprise - The First Adventure",
		  "Cover": 407873,
		  "GenreID": 3,
		  "Price": 137
		},
		{
		  "Title": "Adventures in Consciousness",
		  "Cover": 730805,
		  "GenreID": 4,
		  "Price": 123
		},
		{
		  "Title": "The Secret of the Night: Further Adventures of Rouletabille",
		  "Cover": 2011167,
		  "GenreID": 1,
		  "Price": 152
		},
		{
		  "Title": "Pokémon Adventures, Volume 7",
		  "Cover": 9264303,
		  "GenreID": 2,
		  "Price": 144
		},
		{
		  "Title": "The adventures of Philip on his way through the world",
		  "Cover": 6121766,
		  "GenreID": 4,
		  "Price": 185
		},
		{
		  "Title": "Adventures of Huckleberry Finn / Billy Budd / Red Badge of Courage / Scarlet Letter",
		  "Cover": 8755897,
		  "GenreID": 3,
		  "Price": 96
		},
		{
		  "Title": "The Adventures of Jimmie Dale",
		  "Cover": 2753123,
		  "GenreID": 2,
		  "Price": 153
		},
		{
		  "Title": "The Adventures of Harry Revel",
		  "Cover": 2892232,
		  "GenreID": 3,
		  "Price": 178
		},
		{
		  "Title": "Mermaids Don't Run Track (The Adventures of the Bailey School Kids #26)",
		  "Cover": 6684286,
		  "GenreID": 2,
		  "Price": 54
		},
		{
		  "Title": "Histoires Extraordinaires (Balloon-Hoax / Descent into the Maelstrom / Facts in the Case of M. Valdemar / Gold-Bug / Ligeia / Mesmeric Revelation / Metzengerstein / Morella / MS. Found in a Bottle / Murders in the Rue Morgue / Purloined Letter / Tale of the Ragged Mountains / Unparalleled Adventure of One Hans Pfaall)",
		  "Cover": 11762960,
		  "GenreID": 3,
		  "Price": 187
		},
		{ "Title": "An adventure", "Cover": 7169562, "GenreID": 2, "Price": 91 },
		{
		  "Title": "African adventure",
		  "Cover": 10025744,
		  "GenreID": 3,
		  "Price": 166
		},
		{ "Title": "Gods Adventurer:", "Cover": 1056127, "GenreID": 4, "Price": 106 },
		{
		  "Title": "Adventures in English Literature -- Athena Edition",
		  "Cover": 4938988,
		  "GenreID": 1,
		  "Price": 106
		},
		{
		  "Title": "Sherlock Holmes (Adventures of Sherlock Holmes / Memoirs of Sherlock Holmes [11 stories] / Sign of Four /  Study in Scarlet)",
		  "Cover": 12524376,
		  "GenreID": 1,
		  "Price": 124
		},
		{
		  "Title": "Choose Your Own Adventure - Prisoner of the Ant People",
		  "Cover": 8870239,
		  "GenreID": 2,
		  "Price": 139
		},
		{
		  "Title": "Marvel Adventures Spider-Man",
		  "Cover": 6447846,
		  "GenreID": 1,
		  "Price": 144
		},
		{
		  "Title": "An Awfully Big Adventure",
		  "Cover": 13283869,
		  "GenreID": 3,
		  "Price": 182
		},
		{
		  "Title": "The Adventures of Jerry Muskrat",
		  "Cover": 311420,
		  "GenreID": 1,
		  "Price": 136
		},
		{
		  "Title": "Alice's Adventures in Wonderland [adaptation]",
		  "Cover": 11274845,
		  "GenreID": 1,
		  "Price": 173
		},
		{
		  "Title": "Choose Your Own Adventure - Ghost Hunter",
		  "Cover": 6506412,
		  "GenreID": 2,
		  "Price": 165
		},
		{
		  "Title": "Unbelievably good deals and great adventures that you absolutely can't get unless you're over 50",
		  "Cover": 57934,
		  "GenreID": 2,
		  "Price": 92
		},
		{
		  "Title": "Choose Your Own Adventure - House of Danger",
		  "Cover": 12487964,
		  "GenreID": 4,
		  "Price": 187
		},
		{
		  "Title": "Star Trek Adventures - The Price of The Phoenix",
		  "Cover": 6965676,
		  "GenreID": 4,
		  "Price": 139
		},
		{
		  "Title": "Choose Your Own Adventure - The Lost Jewels of Nabooti",
		  "Cover": 11673617,
		  "GenreID": 4,
		  "Price": 165
		},
		{
		  "Title": "Adventures in two worlds",
		  "Cover": 2516982,
		  "GenreID": 4,
		  "Price": 129
		},
		{
		  "Title": "The Circus of Adventure",
		  "Cover": 11604540,
		  "GenreID": 4,
		  "Price": 140
		},
		{
		  "Title": "The Adventures of Joel Pepper",
		  "Cover": 5907818,
		  "GenreID": 1,
		  "Price": 188
		},
		{ "Title": "Adventure therapy", "Cover": 4736104, "GenreID": 4, "Price": 92 },
		{ "Title": "The adventurers", "Cover": 10233125, "GenreID": 4, "Price": 187 },
		{
		  "Title": "George Washingtons Socks\r\n            \r\n                Time Travel Adventures",
		  "Cover": 7719200,
		  "GenreID": 3,
		  "Price": 114
		},
		{
		  "Title": "Star Trek Adventures - The New Voyages",
		  "Cover": 8666696,
		  "GenreID": 1,
		  "Price": 166
		},
		{
		  "Title": "Wall Street Ventures & Adventures Thru 40 Years",
		  "Cover": 1676776,
		  "GenreID": 3,
		  "Price": 124
		},
		{
		  "Title": "Adventures in Odyssey",
		  "Cover": 1891668,
		  "GenreID": 1,
		  "Price": 70
		},
		{
		  "Title": "Adventures in Singing",
		  "Cover": 1075339,
		  "GenreID": 4,
		  "Price": 116
		},
		{
		  "Title": "Jojo's bizarre adventure",
		  "Cover": 8804928,
		  "GenreID": 1,
		  "Price": 193
		},
		{
		  "Title": "Choose Your Own Adventure - The Horror of High Ridge",
		  "Cover": 368603,
		  "GenreID": 4,
		  "Price": 178
		},
		{
		  "Title": "Escape from the Island of Aquarius (Cooper Kids Adventures",
		  "Cover": 4725299,
		  "GenreID": 4,
		  "Price": 62
		},
		{
		  "Title": "Chilling adventures of Sabrina",
		  "Cover": 8781863,
		  "GenreID": 1,
		  "Price": 142
		},
		{
		  "Title": "Five Fall into Adventure",
		  "Cover": 14571181,
		  "GenreID": 4,
		  "Price": 146
		},
		{
		  "Title": "The adventures of Paddy Beaver",
		  "Cover": 6121764,
		  "GenreID": 2,
		  "Price": 110
		},
		{
		  "Title": "The adventures of Jimmy Skunk",
		  "Cover": 311497,
		  "GenreID": 2,
		  "Price": 200
		},
		{ "Title": "The adventurer", "Cover": 6793323, "GenreID": 1, "Price": 103 },
		{ "Title": "God's adventurer", "Cover": 6505042, "GenreID": 2, "Price": 200 },
		{
		  "Title": "Star Trek Adventures - Spock Must Die!",
		  "Cover": 368623,
		  "GenreID": 2,
		  "Price": 68
		},
		{
		  "Title": "The Incredible Adventure of Professor Branestawm",
		  "Cover": 11955379,
		  "GenreID": 2,
		  "Price": 194
		},
		{
		  "Title": "The Complete Adventures of Peter Rabbit",
		  "Cover": 9044350,
		  "GenreID": 1,
		  "Price": 140
		},
		{
		  "Title": "The Surprising Adventures of Baron Munchausen",
		  "Cover": 9070653,
		  "GenreID": 2,
		  "Price": 121
		},
		{
		  "Title": "The Adventures of Maya the Bee",
		  "Cover": 2788611,
		  "GenreID": 4,
		  "Price": 190
		},
		{
		  "Title": "The adventures of Signor Gaudentio di Lucca",
		  "Cover": 6266798,
		  "GenreID": 3,
		  "Price": 76
		},
		{
		  "Title": "Stanley's Christmas adventure",
		  "Cover": 10647289,
		  "GenreID": 3,
		  "Price": 83
		},
		{
		  "Title": "Surrender in Moonlight:(Love and Adventure Collection#4)",
		  "Cover": 6567326,
		  "GenreID": 4,
		  "Price": 86
		},
		{
		  "Title": "Jamestown, New World adventure",
		  "Cover": 4323389,
		  "GenreID": 4,
		  "Price": 71
		},
		{
		  "Title": "Adventures of George Lee",
		  "Cover": 12040027,
		  "GenreID": 2,
		  "Price": 62
		},
		{
		  "Title": "The Unwanted Undead Adventurer",
		  "Cover": 11615669,
		  "GenreID": 2,
		  "Price": 181
		},
		{
		  "Title": "Choose Your Own Adventure - Superbike",
		  "Cover": 9529341,
		  "GenreID": 4,
		  "Price": 176
		},
		{
		  "Title": "Choose Your Own Adventure - Earthquake!",
		  "Cover": 7074346,
		  "GenreID": 4,
		  "Price": 142
		},
		{
		  "Title": "The adventures of Amelia Bedelia",
		  "Cover": 8009129,
		  "GenreID": 1,
		  "Price": 130
		},
		{ "Title": "Adventure Double", "Cover": 76646, "GenreID": 2, "Price": 79 },
		{
		  "Title": "Brazilian Adventure",
		  "Cover": 595547,
		  "GenreID": 4,
		  "Price": 53
		},
		{
		  "Title": "Choose Your Own Adventure - Mountain Survival",
		  "Cover": 368604,
		  "GenreID": 2,
		  "Price": 143
		},
		{
		  "Title": "The adventures and misadventures of Maqroll",
		  "Cover": 8276735,
		  "GenreID": 2,
		  "Price": 53
		},
		{
		  "Title": "The adventures of Mr. Verdant Green",
		  "Cover": 5756106,
		  "GenreID": 2,
		  "Price": 141
		},
		{
		  "Title": "The Adventure of the Copper Beeches",
		  "Cover": 9246426,
		  "GenreID": 3,
		  "Price": 125
		},
		{
		  "Title": "Star Trek Adventures - Trek to Madworld",
		  "Cover": 368625,
		  "GenreID": 1,
		  "Price": 64
		},
		{
		  "Title": "Narrative of the Life and Adventures of Henry Bibb, An American Slave, Written by Himself",
		  "Cover": 9224679,
		  "GenreID": 3,
		  "Price": 99
		},
		{
		  "Title": "Tenting To-night (A Chronicle of Sport and Adventure in Glacier Park and the Cascade Mountains)",
		  "Cover": 1850617,
		  "GenreID": 1,
		  "Price": 146
		},
		{
		  "Title": "Adventures in friendship",
		  "Cover": 782924,
		  "GenreID": 2,
		  "Price": 196
		},
		{
		  "Title": "The joyous adventure",
		  "Cover": 6519898,
		  "GenreID": 2,
		  "Price": 111
		},
		{
		  "Title": "Choose Your Own Adventure - You Are a Millionaire",
		  "Cover": 6639997,
		  "GenreID": 3,
		  "Price": 150
		},
		{
		  "Title": "Choose Your Own Adventure - Inside UFO 54-40",
		  "Cover": 6979939,
		  "GenreID": 4,
		  "Price": 181
		},
		{
		  "Title": "Choose Your Own Adventure - By Balloon to the Sahara",
		  "Cover": 4672812,
		  "GenreID": 2,
		  "Price": 164
		},
		{
		  "Title": "Walt Disney's The Many Adventures of Winnie the Pooh",
		  "Cover": 10529727,
		  "GenreID": 3,
		  "Price": 184
		},
		{
		  "Title": "Adventures in the Anthropocene",
		  "Cover": 7392616,
		  "GenreID": 4,
		  "Price": 165
		},
		{
		  "Title": "The street of adventure",
		  "Cover": 5623611,
		  "GenreID": 3,
		  "Price": 107
		},
		{
		  "Title": "Adventures Beyond the Body",
		  "Cover": 48669,
		  "GenreID": 3,
		  "Price": 167
		},
		{
		  "Title": "The adventures of the black girl in her search for God",
		  "Cover": 9660488,
		  "GenreID": 1,
		  "Price": 118
		},
		{
		  "Title": "Short Stories (Adventures of Sherlock Holmes / Memoirs of Sherlock Holmes [11 stories])",
		  "Cover": 12504023,
		  "GenreID": 2,
		  "Price": 100
		},
		{
		  "Title": "The Adventures Of Mr. Mocker",
		  "Cover": 763309,
		  "GenreID": 4,
		  "Price": 145
		},
		{
		  "Title": "26 nights, a sexual adventure",
		  "Cover": 6379452,
		  "GenreID": 2,
		  "Price": 156
		},
		{
		  "Title": "Adventures in the skin trade",
		  "Cover": 1534712,
		  "GenreID": 3,
		  "Price": 176
		},
		{ "Title": "The Adventurers", "Cover": 4635795, "GenreID": 1, "Price": 148 },
		{
		  "Title": "Cannibal adventure",
		  "Cover": 6937192,
		  "GenreID": 1,
		  "Price": 56
		},
		{
		  "Title": "Choose Your Own Adventure - Magic Master",
		  "Cover": 11674815,
		  "GenreID": 4,
		  "Price": 144
		},
		{
		  "Title": "The Adventures of Lightfoot the Deer",
		  "Cover": 12173119,
		  "GenreID": 3,
		  "Price": 159
		},
		{
		  "Title": "The adventures of Captain Bonneville, U.S.A. in the Rocky Mountains and the Far West",
		  "Cover": 5555483,
		  "GenreID": 3,
		  "Price": 98
		},
		{
		  "Title": "Choose Your Own Adventure - The Forbidden Castle",
		  "Cover": 11633677,
		  "GenreID": 1,
		  "Price": 197
		},
		{
		  "Title": "Adventures in Contentment",
		  "Cover": 6291667,
		  "GenreID": 1,
		  "Price": 132
		},
		{
		  "Title": "Choose Your Own Adventure - Through the Black Hole",
		  "Cover": 9658345,
		  "GenreID": 4,
		  "Price": 84
		},
		{
		  "Title": "Pokemon. Pathways to Adventure",
		  "Cover": 8783123,
		  "GenreID": 2,
		  "Price": 96
		},
		{
		  "Title": "Real Kids Real Adventures",
		  "Cover": 4628020,
		  "GenreID": 1,
		  "Price": 139
		},
		{
		  "Title": "Star Wars - Han Solo Adventures - Han Solo at Star's End",
		  "Cover": 6624015,
		  "GenreID": 3,
		  "Price": 98
		},
		{
		  "Title": "The Adventures of Hugh Trevor",
		  "Cover": 6121736,
		  "GenreID": 4,
		  "Price": 112
		},
		{
		  "Title": "Adventure of the Empty House",
		  "Cover": 10828471,
		  "GenreID": 3,
		  "Price": 164
		},
		{
		  "Title": "Arthur's Baby (Arthur Adventure Series)",
		  "Cover": 1174971,
		  "GenreID": 1,
		  "Price": 173
		},
		{
		  "Title": "The adventures of Harry Richmond",
		  "Cover": 5813310,
		  "GenreID": 4,
		  "Price": 120
		},
		{
		  "Title": "My adventures as a spy",
		  "Cover": 1861235,
		  "GenreID": 1,
		  "Price": 71
		},
		{
		  "Title": "Arthur's Tooth (Arthur Adventure Series)",
		  "Cover": 188097,
		  "GenreID": 1,
		  "Price": 129
		},
		{
		  "Title": "The Adventures of Peter Rabbit",
		  "Cover": 9660380,
		  "GenreID": 2,
		  "Price": 64
		},
		{ "Title": "Animal Adventures", "Cover": 50661, "GenreID": 3, "Price": 57 },
		{ "Title": "Dolphin Adventure", "Cover": 4090768, "GenreID": 3, "Price": 93 },
		{
		  "Title": "The Further Adventures of Jimmie Dale",
		  "Cover": 1861584,
		  "GenreID": 1,
		  "Price": 84
		},
		{
		  "Title": "Choose Your Own Adventure - Race Forever",
		  "Cover": 8870221,
		  "GenreID": 3,
		  "Price": 155
		},
		{
		  "Title": "Choose Your Own Adventure - Dinosaur Island",
		  "Cover": 8869430,
		  "GenreID": 1,
		  "Price": 192
		},
		{
		  "Title": "The Adventure Zone. Petals to the Metal",
		  "Cover": 8214805,
		  "GenreID": 2,
		  "Price": 111
		},
		{
		  "Title": "The Unlikely Adventures of the Shergill Sisters",
		  "Cover": 8797618,
		  "GenreID": 3,
		  "Price": 51
		},
		{
		  "Title": "What's Your Angle, Pythagoras? A Math Adventure",
		  "Cover": 819019,
		  "GenreID": 4,
		  "Price": 107
		},
		{
		  "Title": "JoJo's Bizarre Adventure",
		  "Cover": 10130815,
		  "GenreID": 4,
		  "Price": 185
		},
		{
		  "Title": "Pirate Island Adventure",
		  "Cover": 282904,
		  "GenreID": 3,
		  "Price": 129
		},
		{
		  "Title": "Choose Your Own Adventure - Lost on the Amazon",
		  "Cover": 11414417,
		  "GenreID": 3,
		  "Price": 72
		},
		{
		  "Title": "Hubert's hair-raising adventure",
		  "Cover": 1529810,
		  "GenreID": 1,
		  "Price": 132
		},
		{
		  "Title": "I Hear Adventure Calling",
		  "Cover": 7285075,
		  "GenreID": 3,
		  "Price": 67
		},
		{
		  "Title": "The Adventures of Pinocchio",
		  "Cover": 11365679,
		  "GenreID": 2,
		  "Price": 197
		},
		{
		  "Title": "The adventures of Ulysses",
		  "Cover": 6379879,
		  "GenreID": 2,
		  "Price": 61
		}
	  ]
	  `

	var books []models.Book
	err := json.Unmarshal([]byte(jsonData), &books)
	if err != nil {
		panic("failed to unmarshal JSON")
	}

	for _, book := range books {
		newBook := models.Book{
			Title:   book.Title,
			Cover:   book.Cover,
			GenreID: book.GenreID,
			Price:   book.Price,
		}

		if err := database.DB.Create(&newBook).Error; err != nil {
			panic("failed to create Book")
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "success create book",
	})
}

func GenreSeeder(c *gin.Context) {
	var genreValue = []string{"Comedy", "Horror", "Adventure", "Thriller"}

	for _, genre := range genreValue {
		newGenre := models.Genre{
			Name: genre,
		}

		if err := database.DB.Create(&newGenre).Error; err != nil {
			panic("failed to create Book")
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "success create genre",
	})
}
