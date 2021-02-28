<?php

namespace classes\handlers;

/**
 * Class Csv
 * @package classes\handlers
 */
class Csv extends Base
{
    public const EXTENSION = 'CSV';

    /**
     * @param string $fileName
     */
    public function handle(string $fileName): void
    {
        if ($this->canHandle($fileName, self::EXTENSION)) {
            echo sprintf('обработчик %s получил файл %s', self::EXTENSION, $fileName) . PHP_EOL;

            $this->writer->Write($this->outFileName, $this->reader->Read($fileName));
        } else {
            parent::handle($fileName);
        }
    }
}
