replicate' :: Int -> Int -> [Int]
replicate' n x  
    | n <= 0    = []  
    | otherwise = x:replicate' (n-1) x  

take' :: (Num i, Ord i) => i -> [a] -> [a]
take' n _
    | n <= 0   = []  
take' _ []     = []  
take' n (x:xs) = x : take' (n-1) xs

maximum' :: (Ord a) => [a] -> a
maximum' [] = error "The list is empty"
maximum' [x] = x
maximum' (x:xs) = max x (maximum' xs)

quicksort :: (Ord a) => [a] -> [a]
quicksort [] = []
quicksort (x:xs) = 
    let left = quicksort [a|a <- xs, a <= x]
        right = quicksort [a|a <- xs, a > x]
    in left ++ [x] ++ right

merge :: (Ord a) => [a] -> [a] -> [a]
merge a [] = a
merge [] a = a
merge (a:as) (b:bs)
    | a < b = a : merge as (b:bs)
    | otherwise = b : merge (a:as) bs

mergesort :: (Ord a) => [a] -> [a]
mergesort [] = []
mergesort [x] = [x]
mergesort arr =
    let halfLength = div (length arr) 2
        left = take halfLength arr
        right = take (length arr - halfLength) $ drop (length left) arr
        sortedLeft = mergesort left
        sortedRight = mergesort right
    in merge sortedLeft sortedRight
    