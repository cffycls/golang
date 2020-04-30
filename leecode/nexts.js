/** 9AM
 * @param {number} n
 * @param {number[][]} relation
 * @param {number} k
 * @return {number}
 */
var numWays = function(n, relation, k) {
    const result = [];
    for (let i = 0; i <= k; i++) {
        result.push(new Array(n).fill(0));
    }
    result[0][0] = 1;
    for (let i = 1; i <=k; i++) {
        for (let j = 0; j < relation.length; j++) {
            result[i][relation[j][1]] += result[i - 1][relation[j][0]];
        }
        console.log(result[i-1])
    }
    console.log(result[k])
    return result[k][n - 1];
};

numWays(5, [[0,2],[2,1],[3,4],[2,3],[1,4],[2,0],[0,4]], 3);

(5) [1, 0, 0, 0, 0]
(5) [0, 0, 1, 0, 1]
(5) [1, 1, 0, 1, 0]
(5) [0, 0, 1, 0, 3]

            0        1        2       3      4        5        6
numWays(4, [[0,2], [2,1], [3,1], [2,3], [1,2], [2,0], [0,2]], 3);
(4) [1, 0, 0, 0]
(4) [0, 0, 2, 0]
                        ret[1][ op[0][1] ] += ret[0][ op[0][0] ] => ret[1][2] = 0+ret[0][0]=1
                        ret[1][ op[1][1] ] += ret[0][ op[1][0] ] => ret[1][1] = 0+ret[0][2]=0
                        ret[1][ op[2][1] ] += ret[0][ op[2][0] ] => ret[1][1] = 0+ret[0][3]=0
                        ret[1][ op[3][1] ] += ret[0][ op[3][0] ] => ret[1][3] = 0+ret[0][2]=0
                        ret[1][ op[4][1] ] += ret[0][ op[4][0] ] => ret[1][2] = 1+ret[0][1]=1
                        ret[1][ op[5][1] ] += ret[0][ op[5][0] ] => ret[1][0] = 0+ret[0][2]=0
                        ret[1][ op[6][1] ] += ret[0][ op[6][0] ] => ret[1][2] = 1+ret[0][0]=2
(4) [2, 2, 0, 2]
                        ret[2][ op[0][1] ] += ret[1][ op[0][0] ] => ret[2][2] = 0+ret[1][0]=0
                        ret[2][ op[1][1] ] += ret[1][ op[1][0] ] => ret[2][1] = 0+ret[1][2]=2
                        ret[2][ op[2][1] ] += ret[1][ op[2][0] ] => ret[2][1] = 2+ret[1][3]=2
                        ret[2][ op[3][1] ] += ret[1][ op[3][0] ] => ret[2][3] = 0+ret[1][2]=2
                        ret[2][ op[4][1] ] += ret[1][ op[4][0] ] => ret[2][2] = 0+ret[1][1]=0
                        ret[2][ op[5][1] ] += ret[1][ op[5][0] ] => ret[2][0] = 0+ret[1][2]=2
                        ret[2][ op[6][1] ] += ret[1][ op[6][0] ] => ret[2][2] = 0+ret[1][0]=0
(4) [0, 2, 6, 0]
                        ret[3][ op[0][1] ] += ret[2][ op[0][0] ] => ret[3][2] = 0+ret[2][0]=2
                        ret[3][ op[1][1] ] += ret[2][ op[1][0] ] => ret[3][1] = 0+ret[2][2]=0
                        ret[3][ op[2][1] ] += ret[2][ op[2][0] ] => ret[3][1] = 0+ret[2][3]=2
                        ret[3][ op[3][1] ] += ret[2][ op[3][0] ] => ret[3][3] = 0+ret[2][2]=0
                        ret[3][ op[4][1] ] += ret[2][ op[4][0] ] => ret[3][2] = 2+ret[2][1]=4
                        ret[3][ op[5][1] ] += ret[2][ op[5][0] ] => ret[3][0] = 0+ret[2][2]=0
                        ret[3][ op[6][1] ] += ret[2][ op[6][0] ] => ret[3][2] = 4+ret[2][0]=6