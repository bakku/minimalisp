(defvar names '(1 2 3 4))
(defun quarter-of (n)
  (let (half (/ n 2))
    (/ half 2)))
  
(let (quarters (map quarter-of names))
  (if (= 0.25 (first quarters))
    (println "yes!")
    (println "no!")))