package controllers

const EMOJI_CAT_FACE = string(128049)
const EMOJI_CAT_BODY = string(128008)
const EMOJI_CAT_GRINNING = string(128568)
const EMOJI_CAT_LAUGHING = string(128569)
const EMOJI_CAT_SMILING = string(128570)
const EMOJI_CAT_HEART_EYES = string(128571)
const EMOJI_CAT_WRY_SMILE = string(128572)
const EMOJI_CAT_KISSING = string(128573)
const EMOJI_CAT_ANGRY = string(128574)
const EMOJI_CAT_CRYING = string(128575)
const EMOJI_CAT_SCARED = string(128576)

const MESSAGE_API_PROBLEM = "I'm sorry I cannot handle your request right now. There is some problem with our questions API. " + EMOJI_CAT_SCARED
const MESSAGE_WANT_QUESTION = "Meow meow! " + EMOJI_CAT_SMILING + " Do you want me to ask you a question?\n\nA ) Yes\nB ) No"
const MESSAGE_BYE = "Bye! Come back again soon! " + EMOJI_CAT_KISSING
const MESSAGE_INVALID_ANSWER = "That is not a valid response. " + EMOJI_CAT_ANGRY + " Please choose a letter from A-%c."
const MESSAGE_CORRECT_ANSWER = "That's right! Meow meow! " + EMOJI_CAT_HEART_EYES + "\n%s is the correct answer!\nYour score increased by 2 points.\n\nDo you want me to ask you another question?\n\nA ) Yes\nB ) No"
const MESSAGE_INCORRECT_ANSWER = "That's incorrect! Meeooww. " + EMOJI_CAT_WRY_SMILE + "\n%s is the correct answer!\nYour score decreased by 1 point. " + EMOJI_CAT_CRYING + "\n\nDo you want me to ask you another question?\n\nA ) Yes\nB ) No"
const MESSAGE_POSITIVE_SCORE = "Your current score is %d. Nice! " + EMOJI_CAT_GRINNING
const MESSAGE_NEGATIVE_SCORE = "Your current score is %d. LMAO noob! " + EMOJI_CAT_LAUGHING
