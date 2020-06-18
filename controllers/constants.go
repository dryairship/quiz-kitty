package controllers

var EMOJI_CAT_FACE = string(128049)
var EMOJI_CAT_BODY = string(128008)
var EMOJI_CAT_GRINNING = string(128568)
var EMOJI_CAT_LAUGHING = string(128569)
var EMOJI_CAT_SMILING = string(128570)
var EMOJI_CAT_HEART_EYES = string(128571)
var EMOJI_CAT_WRY_SMILE = string(128572)
var EMOJI_CAT_KISSING = string(128573)
var EMOJI_CAT_ANGRY = string(128574)
var EMOJI_CAT_CRYING = string(128575)
var EMOJI_CAT_SCARED = string(128576)

var MESSAGE_API_PROBLEM = "I'm sorry I cannot handle your request right now. There is some problem with our questions API. " + EMOJI_CAT_SCARED
var MESSAGE_WANT_QUESTION = "Meow meow! " + EMOJI_CAT_SMILING + " Do you want me to ask you a question?\n\nA ) Yes\nB ) No"
var MESSAGE_BYE = "Bye! Come back again soon! " + EMOJI_CAT_KISSING
var MESSAGE_INVALID_ANSWER = "That is not a valid response. " + EMOJI_CAT_ANGRY + " Please choose a letter from A-%c."
var MESSAGE_CORRECT_ANSWER = "That's right! Meow meow! " + EMOJI_CAT_HEART_EYES + "\n%s is the correct answer!\n\nDo you want me to ask you another question?\n\nA ) Yes\nB ) No"
var MESSAGE_INCORRECT_ANSWER = "That's incorrect! Meeooww. " + EMOJI_CAT_CRYING + "\n%s is the correct answer!\n\nDo you want me to ask you another question?\n\nA ) Yes\nB ) No"
var MESSAGE_POSITIVE_SCORE = "Your current score is %d. Nice! " + EMOJI_CAT_GRINNING
var MESSAGE_NEGATIVE_SCORE = "Your current score is %d. LMAO noob! " + EMOJI_CAT_LAUGHING
