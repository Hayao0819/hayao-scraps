#!/usr/bin/env runghc

import Distribution.Simple.Utils (xargs)
data FizzBuzz = Fizz | Buzz | FizzBuzz | FBInt Int

instance Show FizzBuzz where
    show n = case n of
        FizzBuzz -> "FizzBuzz"
        Fizz -> "Fizz"
        Buzz -> "Buzz"
        FBInt x -> show x

determineFZ :: Int -> FizzBuzz
determineFZ x
  | x `mod` 15 == 0 = FizzBuzz
  | x `mod` 3 == 0 = Fizz
  | x `mod` 5 == 0 = Buzz
  | otherwise  = FBInt x

fizzBuzzLoop :: Int -> IO ()
fizzBuzzLoop x = do
    print (determineFZ x)
    fizzBuzzLoop(x+1)

main = fizzBuzzLoop 1
