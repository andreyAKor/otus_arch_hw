<?php
require_once __DIR__ . '/vendor/autoload.php';

use classes\matrix\Matrix;
use classes\matrix\MatrixAdditionAdapter;
use classes\matrix\MatrixGenarator;
use classes\writer\WriteFileMatrixes;

if (count($argv) != 3) {
    echo 'Неверный формат запуска. Пример: php ' . $argv[0] . ' F2.txt F1.txt' . PHP_EOL;
    exit(1);
}

unset($argv[0]);

try {
    run(...$argv);
} catch (Exception $e) {
    echo 'Исключение: ' . $e->getMessage() . PHP_EOL;
    exit(1);
}

function run(string $outFileMatrixGenerator, string $outFileMatrixAddition)
{
    (new MatrixGenarator(new WriteFileMatrixes($outFileMatrixGenerator)))->Run();

    $operation = new MatrixAdditionAdapter($outFileMatrixGenerator, $outFileMatrixAddition);
    (new Matrix($operation))->Calc();
}
