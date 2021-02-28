<?php

namespace classes\handlers;

/**
 * Class Txt
 * @package classes\handlers
 */
class Txt extends Base
{
    public const EXTENSION = 'TXT';

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
